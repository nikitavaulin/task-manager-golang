package core_http_middleware

import (
	"fmt"
	"net/http"

	tools_envparser "github.com/nikitavaulin/task-manager-golang/internal/core/tools/env_parser"
	tools_jwt "github.com/nikitavaulin/task-manager-golang/internal/core/tools/jwt"
	tools_passwordhasher "github.com/nikitavaulin/task-manager-golang/internal/core/tools/password_hasher"
)

func Auth() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			appPassword, err := tools_envparser.GetAppPassword()
			if err != nil {
				errMsg := fmt.Sprintf("ERROR: failed to get app password: %v\n", err)
				http.Error(w, errMsg, http.StatusInternalServerError)
				return
			}

			var jwt string
			cookie, err := r.Cookie("token")
			if err == nil {
				jwt = cookie.Value
			}

			if len(jwt) == 0 {
				http.Error(w, "Authentification required", http.StatusUnauthorized)
				return
			}
			passwordHash, ok := getUserPasswordHashFromJWT(jwt)
			if !ok {
				http.Error(w, "Authentification required", http.StatusUnauthorized)
				return
			}

			if !tools_passwordhasher.VerifyPassword(appPassword, passwordHash) {
				http.Error(w, "Authentification required", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func getUserPasswordHashFromJWT(jwt string) (string, bool) {
	claims, err := tools_jwt.DecodeClaims(jwt)
	if err != nil {
		return "", false
	}

	value, ok := claims["password"]
	if !ok {
		return "", false
	}

	passwordHash := fmt.Sprint(value)
	return passwordHash, true

}
