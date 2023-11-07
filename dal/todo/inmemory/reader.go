package inmemory

import (
	"fmt"

	"github.com/hobord/poc-htmx-go-todolist/entities"
)

func (d *dal) GetByID(id string) (*entities.Todo, error) {
	d.lck.RLock()
	defer d.lck.RUnlock()

	for _, group := range d.data {
		for _, todos := range group {
			for _, todo := range todos {
				if todo.ID == id {
					return todo, nil
				}
			}
		}
	}

	return nil, fmt.Errorf("todo not found")
}

func (d *dal) GetAll(user string) ([]*entities.Todo, error) {
	d.lck.RLock()
	defer d.lck.RUnlock()

	todos := make([]*entities.Todo, 0, 10)
	for _, group := range d.data[user] {
		todos = append(todos, group...)
	}

	return todos, nil
}

func (d *dal) GetByGroup(user string, group string) ([]*entities.Todo, error) {
	d.lck.RLock()
	defer d.lck.RUnlock()

	return d.data[user][group], nil
}
