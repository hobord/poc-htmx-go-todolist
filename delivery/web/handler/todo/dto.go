package todo

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type CreateTodoGroupRequest struct {
	Title string `json:"title"`
	Color string `json:"color,omitempty"`
}

func (dto *CreateTodoGroupRequest) Bind(r *http.Request) error {
	dto.Title = r.FormValue("title")
	if dto.Title == "" {
		return fmt.Errorf("name is required")
	}

	dto.Color = url.QueryEscape(r.FormValue("color"))

	return nil
}

type UpdateTodoGroupRequest struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Color string `json:"color,omitempty"`
}

func (dto *UpdateTodoGroupRequest) Bind(r *http.Request) error {
	dto.ID = r.FormValue("id")
	if dto.ID == "" {
		return fmt.Errorf("id is required")
	}

	dto.Title = r.FormValue("name")
	if dto.Title == "" {
		return fmt.Errorf("name is required")
	}

	dto.Color = url.QueryEscape(r.FormValue("color"))

	return nil
}

type CreateTodoItemRequest struct {
	GroupID  string `json:"group_id"`
	Title    string `json:"title"`
	Priority int    `json:"priority,omitempty"`
}

func (dto *CreateTodoItemRequest) Bind(r *http.Request) error {
	dto.GroupID = r.FormValue("group_id")
	if dto.GroupID == "" {
		return fmt.Errorf("group_id is required")
	}

	dto.Title = r.FormValue("title")
	if dto.Title == "" {
		return fmt.Errorf("title is required")
	}

	dto.Priority, _ = strconv.Atoi(r.FormValue("priority"))

	return nil
}

type UpdateTodoItemRequest struct {
	ID        string `json:"id"`
	GroupID   string `json:"group_id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	Priority  int    `json:"priority,omitempty"`
}

func (dto *UpdateTodoItemRequest) Bind(r *http.Request) error {
	dto.ID = r.FormValue("id")
	if dto.ID == "" {
		return fmt.Errorf("id is required")
	}

	dto.GroupID = r.FormValue("group_id")
	if dto.GroupID == "" {
		return fmt.Errorf("group_id is required")
	}

	dto.Title = r.FormValue("title")
	if dto.Title == "" {
		return fmt.Errorf("title is required")
	}

	dto.Priority, _ = strconv.Atoi(r.FormValue("priority"))

	dto.Completed = r.FormValue("completed") == "true"

	return nil
}

type SortItemsRequest struct {
	Items []string `json:"items"`
}

func (dto *SortItemsRequest) Bind(r *http.Request) error {
	r.ParseForm()

	dto.Items = r.Form["item"]

	return nil
}
