package api

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/hardikroongta8/choplinks/auth"
	"github.com/hardikroongta8/choplinks/model"
	"net/http"
)

func (s *Server) handleUserRoutes(r *mux.Router) {
	r.Handle("/", errorHandler(s.handleCreateUser)).Methods("POST")
	r.Handle("/", r.MethodNotAllowedHandler)
}

func (s *Server) handleCreateUser(w http.ResponseWriter, r *http.Request) error {
	createUserBody := new(model.CreateUserReqBody)
	if err := json.NewDecoder(r.Body).Decode(createUserBody); err != nil {
		return err
	}
	newUser := model.User{
		ID:    uuid.NewString(),
		Name:  createUserBody.Name,
		Email: createUserBody.Email,
	}
	token, err := auth.GenerateJWT(newUser.ID)
	if err != nil {
		return err
	}
	err = s.store.CreateUser(&newUser)
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusCreated, apiResponse{Data: map[string]string{
		"userID":      newUser.ID,
		"accessToken": token,
	}})
}
