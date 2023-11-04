package health

import (
	"fmt"
	"net/http"
)

//go:generate mockery --name Handler --structname MockHandler --output . --outpkg health --case underscore --filename handler_mock.go
type Handler interface {
	Health(w http.ResponseWriter, r *http.Request)
}

type handler struct {
}

func NewCheck() Handler {
	return &handler{}
}

func (h *handler) Health(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "ok")
}
