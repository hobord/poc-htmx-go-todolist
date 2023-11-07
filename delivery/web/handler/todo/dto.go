package todo

import (
	"fmt"
	"net/http"
)

type AddItemRequest struct {
	Title string `json:"title"`
	Group string `json:"group"`
}

func (d *AddItemRequest) ParseRequest(r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	d.Title = r.FormValue("title")
	d.Group = r.FormValue("group")

	return d.Validate()
}

func (d *AddItemRequest) Validate() error {
	if d.Title == "" {
		return fmt.Errorf("title is empty")
	}

	if d.Group == "" {
		return fmt.Errorf("group is empty")
	}

	return nil
}
