package core_http_middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nikitavaulin/task-manager-golang/internal/core/domain"
	tools_jwt "github.com/nikitavaulin/task-manager-golang/internal/core/tools/jwt"
)

func Auth() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var jwt string
			cookie, err := r.Cookie("token")
			if err == nil {
				jwt = cookie.Value
			}

			if len(jwt) == 0 {
				http.Error(w, "Authentification required (empty cookie)", http.StatusUnauthorized)
				return
			}

			username, ok := getUsernameFromJWT(jwt)
			if !ok {
				http.Error(w, "Authentification required (invalid token)", http.StatusUnauthorized)
				return
			}

			ctx := r.Context()
			ctx = domain.UsernameToContext(ctx, username)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func getUsernameFromJWT(jwt string) (string, bool) {
	claims, err := tools_jwt.DecodeClaims(jwt)
	if err != nil {
		log.Printf("ERROR: %v", err)
		return "", false
	}

	value, ok := claims["username"]
	if !ok {
		return "", false
	}

	username := fmt.Sprint(value)
	return username, true

}
