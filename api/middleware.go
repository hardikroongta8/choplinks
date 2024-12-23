package api

import (
	"context"
	"github.com/hardikroongta8/choplinks/auth"
	"net/http"
	"strings"
)

func authMiddleware(next http.Handler) http.Handler {
	return errorHandler(func(w http.ResponseWriter, r *http.Request) error {
		h := r.Header.Get("Authorization")
		if len(h) == 0 {
			return writeJSON(w, http.StatusBadRequest, "authentication header missing")
		}
		arr := strings.Split(h, " ")
		if len(arr) != 2 {
			return writeJSON(w, http.StatusBadRequest, "invalid token format")
		}
		if arr[0] != "Bearer" {
			return writeJSON(w, http.StatusBadRequest, "invalid token format")
		}
		userID, err := auth.ParseJWT(arr[1])
		if err != nil {
			return writeJSON(w, http.StatusUnauthorized, err.Error())
		}
		ctx := context.WithValue(r.Context(), "userID", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
		return nil
	})
}
