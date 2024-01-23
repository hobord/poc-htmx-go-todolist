package todo

import (
	"fmt"

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
