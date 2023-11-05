package index

import (
	"net/http"

	"github.com/hobord/poc-htmx-go-todolist/delivery/web/templates/layouts"
	"github.com/hobord/poc-htmx-go-todolist/services/todo"
)

type handler struct {
	todoService todo.Service
}

func NewHandler(todoService todo.Service) Handler {
	return &handler{
		todoService: todoService,
	}
}

func (h *handler) IndexPage(w http.ResponseWriter, r *http.Request) {
	todos, err := h.todoService.GetAllGroup("user1")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	layout := layouts.IndexPage(todos)

	if err := layout.Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
