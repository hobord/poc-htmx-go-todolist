package index

import (
	"net/http"

	"github.com/hobord/poc-htmx-go-todolist/delivery/web/templates/layouts"
)

type handler struct {
}

func NewHandler() Handler {
	return &handler{}
}

func (h *handler) IndexPage(w http.ResponseWriter, r *http.Request) {
	layout := layouts.IndexPage()

	layout.Render(r.Context(), w)
}
