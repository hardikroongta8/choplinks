package api

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/hardikroongta8/choplinks/config"
	"github.com/hardikroongta8/choplinks/model"
	"math/rand"
	"net/http"
	"time"
)

func (s *Server) handleURLMapRoutes(r *mux.Router) {
	r.Use(authMiddleware)
	r.Handle("/", errorHandler(s.handleCreateURLMap)).Methods("POST")
	r.Handle("/", errorHandler(s.handleGetURLMaps)).Methods("GET")
	r.Handle("/", errorHandler(s.handleDeleteURLMap)).Methods("DELETE")

	r.Handle("/", r.MethodNotAllowedHandler)
}

func (s *Server) handleCreateURLMap(w http.ResponseWriter, r *http.Request) error {
	createUrlMapBody := new(model.CreateURLMapReqBody)
	if err := json.NewDecoder(r.Body).Decode(createUrlMapBody); err != nil {
		return err
	}
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	urlMap := model.URLMap{
		ID:          string(b),
		OriginalURL: createUrlMapBody.OriginalUrl,
		UserID:      r.Context().Value("userID").(string),
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
	err := s.store.CreateURLMap(&urlMap)
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusCreated, apiResponse{
		Data: config.Get().BaseURL + "/" + urlMap.ID,
	})
}

func (s *Server) handleGetURLMaps(w http.ResponseWriter, r *http.Request) error {
	userID := r.Context().Value("userID").(string)
	urlMaps, err := s.store.GetAllURLMapsByUserID(userID)
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, apiResponse{Data: urlMaps})
}

func (s *Server) handleDeleteURLMap(w http.ResponseWriter, r *http.Request) error {
	userID := r.Context().Value("userID").(string)
	id := r.URL.Query().Get("id")
	urlMap, err := s.store.GetURLMapByID(id)
	if err != nil {
		return err
	}
	if urlMap.UserID != userID {
		return errors.New("requested resource does not exist")
	}
	err = s.store.DeleteURLMapByID(id)
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, apiResponse{Data: id})
}
