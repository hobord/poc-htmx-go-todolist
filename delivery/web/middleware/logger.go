package middleware

import (
	"fmt"
	"net/http"

	"github.com/hobord/poc-htmx-go-todolist/delivery/web/router"
	"github.com/hobord/poc-htmx-go-todolist/services/logger"
)

func Logger(next http.Handler) http.Handler {
	return router.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request: %s %s\n", r.Method, r.URL.Path)
		// do stuff
		next.ServeHTTP(w, r)
	})
}

func WithLogger(log logger.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return router.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Debug("HttpRequest", "method", r.Method, "path", r.URL.Path)
			// do stuff
			next.ServeHTTP(w, r)
		})
	}
}
