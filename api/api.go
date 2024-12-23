package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/hardikroongta8/choplinks/storage"
	"log"
	"net/http"
)

type Server struct {
	listenAddr string
	store      storage.Store
}

func NewServer(listenAddr string, store storage.Store) *Server {
	return &Server{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *Server) Run() {
	r := mux.NewRouter()
	r.StrictSlash(true)
	s.handleURLMapRoutes(r.PathPrefix("/url").Subrouter())
	s.handleUserRoutes(r.PathPrefix("/user").Subrouter())
	r.Handle("/{id}", errorHandler(s.handleRedirect))

	log.Println("Server running on port:", s.listenAddr)
	log.Fatalln(http.ListenAndServe(s.listenAddr, r))
}

func (s *Server) handleRedirect(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return errors.New("ID param is missing")
	}
	urlMap, err := s.store.GetURLMapByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return writeJSON(w, http.StatusNotFound, apiResponse{Error: "Invalid URL"})
		}
		return err
	}
	http.Redirect(w, r, urlMap.OriginalURL, http.StatusFound)
	return nil
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type apiResponse struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func errorHandler(f apiFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			err = writeJSON(w, http.StatusBadRequest, apiResponse{Error: err.Error()})
			if err != nil {
				log.Println("Error writing response:", err.Error())
			}
		}
	})
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}
