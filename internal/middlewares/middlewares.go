package middlewares

import (
	"errors"
	"net/http"

	"github.com/Raybird/whale/internal/auth"
	"github.com/Raybird/whale/internal/responses"
)

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		next(w, r)
	}
}
