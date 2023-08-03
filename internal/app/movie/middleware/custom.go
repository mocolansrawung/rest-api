package middleware

import (
	"log"
	"net/http"
)

func CustomMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Custom middleware is executed")

		next.ServeHTTP(w, r)
	})
}
