package middlewares

import (
	"net/http"
	"todo/core/response"
	"todo/utils/jwt"

	"github.com/gorilla/context"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Actions before controller executed
		jwtToken := jwt.ParseToken(r.Header.Get("authorization"))
		if jwtToken["status"] == "errors" {
			response.JSON(w, http.StatusUnauthorized, jwtToken)
			return
		}

		context.Set(r, "user_id", jwtToken["user_id"])

		// Execute Controller
		next.ServeHTTP(w, r)

		// Actions after controller executed
	})
}
