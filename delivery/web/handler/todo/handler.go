package todo

import (
	"fmt"
	"net/http"

	"github.com/hobord/poc-htmx-go-todolist/delivery/web/templates/components"
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

func (h *handler) AddItem(w http.ResponseWriter, r *http.Request) {
	dto := &AddItemRequest{}

	err := dto.ParseRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, err := h.todoService.Create("user1", dto.Group, dto.Title); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	todos, err := h.todoService.GetByGroup("user1", "ds")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var title string
	if len(todos) > 0 {
		title = todos[0].Group
	}

	content := components.TodoList(title, todos)

	if err := content.Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func (h *handler) TodoLists(w http.ResponseWriter, r *http.Request) {

}

func (h *handler) TodoItems(w http.ResponseWriter, r *http.Request) {

}

func (h *handler) TodoItem(w http.ResponseWriter, r *http.Request) {

}
