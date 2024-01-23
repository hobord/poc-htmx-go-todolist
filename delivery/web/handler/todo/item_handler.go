package todo

import (
	"fmt"
	"net/http"

	"github.com/hobord/poc-htmx-go-todolist/delivery/web/router"
	"github.com/hobord/poc-htmx-go-todolist/delivery/web/templates/components"
	"github.com/hobord/poc-htmx-go-todolist/entities"
)

func (h *handler) CreateItem(w http.ResponseWriter, r *http.Request) {
	var dto CreateTodoItemRequest

	if err := dto.Bind(r); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	item := &entities.TodoItem{
		GroupID:  dto.GroupID,
		Title:    dto.Title,
		Priority: dto.Priority,
	}

	if err := h.todoService.AddTodoItem(item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	group, err := h.todoService.GetTodoGroupByID(dto.GroupID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	component := components.TodoGroup(group)

	if err := component.Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func (h *handler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	var dto UpdateTodoItemRequest

	if err := dto.Bind(r); err != nil {

		return
	}

	item := &entities.TodoItem{
		ID:        dto.ID,
		GroupID:   dto.GroupID,
		Title:     dto.Title,
		Completed: dto.Completed,
		Priority:  dto.Priority,
	}

	if err := h.todoService.UpdateTodoItem(item); err != nil {
		return
	}

	fmt.Fprintf(w, "%+v", item)
}

func (h *handler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	groupID := router.ParamsFromURL(r, "groupID")
	if groupID == "" {
		return
	}

	itemID := router.ParamsFromURL(r, "itemID")
	if itemID == "" {
		return
	}

	if err := h.todoService.DeleteTodoItem(itemID); err != nil {
		return
	}
}

func (h *handler) DeleteCompletedItems(w http.ResponseWriter, r *http.Request) {
	groupID := router.ParamsFromURL(r, "groupID")
	if groupID == "" {
		return
	}

	if err := h.todoService.DeleteCompletedTodoItems(groupID); err != nil {
		return
	}

	group, err := h.todoService.GetTodoGroupByID(groupID)
	if err != nil {
		return
	}

	fmt.Fprintf(w, "%+v", group)
}
