package inmemory

import (
	"github.com/google/uuid"
	"github.com/hobord/poc-htmx-go-todolist/entities"
)

func generateID() string {
	id := uuid.New()
	return id.String()
}

func (d *dal) WriteTodoGroup(todoGroup *entities.TodoGroup) error {
	d.lck.Lock()
	defer d.lck.Unlock()

	if todoGroup.ID == "" {
		todoGroup.ID = generateID()
	}

	dto := toGroupDTO(todoGroup)

	if err := d.writeTodoGroup(dto); err != nil {
		return err
	}

	if len(todoGroup.Todos) > 0 {
		d.wirteTodoItems(todoGroup.Todos)
	}

	return nil
}

func (d *dal) writeTodoGroup(dto groupDTO) error {
	d.groups[dto.ID] = dto

	return nil
}

func (d *dal) DeleteTodoGroup(todoGroupID string) error {
	d.lck.Lock()
	defer d.lck.Unlock()

	delete(d.groups, todoGroupID)

	return d.deleteTodoItemsByGroupID(todoGroupID)
}

func (d *dal) deleteTodoItemsByGroupID(todoGroupID string) error {
	for i, item := range d.items {
		if item.GroupID == todoGroupID {
			delete(d.items, i)
		}
	}

	return nil
}

func (d *dal) WriteTodoItem(todoItem *entities.TodoItem) error {
	d.lck.Lock()
	defer d.lck.Unlock()

	if todoItem.ID == "" {
		todoItem.ID = generateID()
	}

	dto := toItemDTO(todoItem)

	return d.writeItem(dto)
}

func (d *dal) wirteTodoItems(todoItems []*entities.TodoItem) error {
	for _, todoItem := range todoItems {
		item := toItemDTO(todoItem)
		if err := d.writeItem(item); err != nil {
			return err
		}
	}

	return nil
}

func (d *dal) writeItem(dto itemDTO) error {
	d.items[dto.ID] = dto

	return nil
}

func (d *dal) DeleteTodoItem(todoItemID string) error {
	d.lck.Lock()
	defer d.lck.Unlock()

	delete(d.items, todoItemID)

	return nil
}
