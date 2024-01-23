package todo

import (
	"fmt"
	"net/http"

	"github.com/hobord/poc-htmx-go-todolist/delivery/web/router"
	"github.com/hobord/poc-htmx-go-todolist/delivery/web/templates/components"
	"github.com/hobord/poc-htmx-go-todolist/entities"
)

func (h *handler) CreateTodoGroup(w http.ResponseWriter, r *http.Request) {
	var dto CreateTodoGroupRequest

	if err := dto.Bind(r); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	group := &entities.TodoGroup{
		Title:  dto.Title,
		Color:  dto.Color,
		UserID: "1",
	}

	if err := h.todoService.CreateTodoGroup(group); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	groups, err := h.todoService.GetTodoGroupsByUserID("1")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	component := components.ListTodoGroups(groups)

	if err := component.Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *handler) GetTodoGroup(w http.ResponseWriter, r *http.Request) {
	groupID := router.ParamsFromURL(r, "groupID")

	group, err := h.todoService.GetTodoGroupByID(groupID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%+v", group)
}

func (h *handler) GetTodoGroups(w http.ResponseWriter, r *http.Request) {
	groups, err := h.todoService.GetTodoGroupsByUserID("1")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%+v", groups)
}

func (h *handler) UpdateTodoGroup(w http.ResponseWriter, r *http.Request) {
	groupID := router.ParamsFromURL(r, "groupID")

	group, err := h.todoService.GetTodoGroupByID(groupID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var dto UpdateTodoGroupRequest

	if err := dto.Bind(r); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.todoService.UpdateTodoGroup(group); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%+v", group)
}

func (h *handler) DeleteTodoGroup(w http.ResponseWriter, r *http.Request) {
	groupID := router.ParamsFromURL(r, "groupID")

	err := h.todoService.DeleteTodoGroup(groupID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// groups, err := h.todoService.GetTodoGroupsByUserID("1")
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	//
	// component := components.ListTodoGroups(groups)
	//
	// if err := component.Render(r.Context(), w); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }
}
