package index

import (
	"net/http"

	"github.com/hobord/poc-htmx-go-todolist/delivery/web/templates/layouts"
	"github.com/hobord/poc-htmx-go-todolist/entities"
)

type handler struct {
}

func NewHandler() Handler {
	return &handler{}
}

func (h *handler) IndexPage(w http.ResponseWriter, r *http.Request) {
	var todos map[string][]*entities.Todo

	// TODO: get todos from database

	layout := layouts.IndexPage(todos)

	if err := layout.Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
