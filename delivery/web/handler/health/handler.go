package health

import (
	"fmt"
	"net/http"

	"github.com/hobord/poc-htmx-go-todolist/services/health"
)

//go:generate mockery --name Handler --structname MockHandler --output . --outpkg health --case underscore --filename handler_mock.go
type Handler interface {
	Health(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	service health.Service
}

func NewCheck(service health.Service) (Handler, error) {
	h := &handler{
		service: service,
	}

	return h, h.validate()
}

func (h *handler) validate() error {
	if h.service == nil {
		return fmt.Errorf("service is nil")
	}

	return nil
}

func (h *handler) Health(w http.ResponseWriter, _ *http.Request) {
	if err := h.service.Health(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error: %v", err)
		return
	}

	fmt.Fprintf(w, "ok")
}
