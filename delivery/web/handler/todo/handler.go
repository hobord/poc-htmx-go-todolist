package todo

import (
	"fmt"
	"net/http"

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

func (h *handler) TodoLists(w http.ResponseWriter, r *http.Request) {

}

func (h *handler) TodoItems(w http.ResponseWriter, r *http.Request) {

}
func (h *handler) TodoItem(w http.ResponseWriter, r *http.Request) {

}
