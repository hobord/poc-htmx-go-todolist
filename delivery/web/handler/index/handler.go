package index

import (
	"fmt"
	"net/http"

	"github.com/hobord/poc-htmx-go-todolist/delivery/web/templates/views"
	"github.com/hobord/poc-htmx-go-todolist/services/todo"
)

type handler struct {
	todoService todo.Service
}

func NewHandler(todoService todo.Service) (Handler, error) {
	h := &handler{
		todoService: todoService,
	}

	return h, h.validate()
}

func (h *handler) validate() error {
	if h.todoService == nil {
		return fmt.Errorf("todoService is nil")
	}

	return nil
}

func (h *handler) IndexPage(w http.ResponseWriter, r *http.Request) {
	groups, err := h.todoService.GetTodoGroupsByUserID("")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	view := views.IndexPage(groups)

	if err := view.Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
