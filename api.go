package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type APIServer struct {
	listenAddr string
	store      Storage
}

func NewAPIServer(listenAddr string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/user", errorHandler(s.handleUser))
	router.HandleFunc("/url", errorHandler(s.handleURLMap)) // ADD JWT MIDDLEWARE
	//router.Handle("/user", http.NewServeMux())

	log.Println("Server running on port:", s.listenAddr)
	log.Fatalln(http.ListenAndServe(s.listenAddr, router))
}

func (s *APIServer) handleUser(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		return s.handleCreateUser(w, r)
	}
	return fmt.Errorf("method not allowed: %s", r.Method)
}

func (s *APIServer) handleURLMap(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		return s.handleCreateURLMap(w, r)
	}
	if r.Method == "GET" {
		return s.handleGetURLMaps(w, r)
	}
	return fmt.Errorf("method not allowed: %s", r.Method)
}

func (s *APIServer) handleCreateUser(w http.ResponseWriter, r *http.Request) error {
	createUserBody := new(CreateUserReqBody)
	if err := json.NewDecoder(r.Body).Decode(createUserBody); err != nil {
		return err
	}
	newUser := User{
		ID:    uuid.NewString(),
		Name:  createUserBody.Name,
		Email: createUserBody.Email,
	}
	err := s.store.CreateUser(&newUser)
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusCreated, apiResponse{Data: newUser.ID})
}

func (s *APIServer) handleCreateURLMap(w http.ResponseWriter, r *http.Request) error {
	createUrlMapBody := new(CreateURLMapReqBody)
	if err := json.NewDecoder(r.Body).Decode(createUrlMapBody); err != nil {
		return err
	}
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	urlMap := URLMap{
		ShortenedURLPath: string(b),
		OriginalURL:      createUrlMapBody.OriginalUrl,
		UserID:           createUrlMapBody.UserID,
		CreatedAt:        time.Now().UTC(),
		UpdatedAt:        time.Now().UTC(),
	}
	err := s.store.CreateURLMap(&urlMap)
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusCreated, apiResponse{
		Data: GetConfig().Server.BaseURL + "/" + urlMap.ShortenedURLPath,
	})
}

func (s *APIServer) handleGetURLMaps(w http.ResponseWriter, r *http.Request) error {
	getUrlMapsBody := new(GetURLMapsReqBody)
	if err := json.NewDecoder(r.Body).Decode(getUrlMapsBody); err != nil {
		return err
	}
	urlMaps, err := s.store.GetAllURLMapsByUserID(getUrlMapsBody.UserID)
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, apiResponse{Data: urlMaps})
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type apiResponse struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func errorHandler(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			writeJSON(w, http.StatusBadRequest, apiResponse{Error: err.Error()})
		}
	}
}
