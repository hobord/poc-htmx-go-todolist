package middleware

import (
	"fmt"
	"net/http"

	"github.com/hobord/poc-htmx-go-todolist/delivery/web/router"
)

func Logger(next http.Handler) http.Handler {
	return router.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request: %s %s\n", r.Method, r.URL.Path)
		// do stuff
		next.ServeHTTP(w, r)
	})
}
