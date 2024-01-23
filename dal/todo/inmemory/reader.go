package inmemory

import (
	"github.com/hobord/poc-htmx-go-todolist/entities"
)

func (d *dal) GetTodoGroupsByUserID(userID string) ([]*entities.TodoGroup, error) {
	d.lck.RLocker().Lock()
	defer d.lck.RLocker().Unlock()

	var todoGroups []*entities.TodoGroup

	for _, g := range d.groups {
		if g.UserID == userID {
			todoGroup := g.toGroupEntity()
			todoGroup.Todos = d.getItemsByGroupID(g.ID)

			todoGroups = append(todoGroups, todoGroup)
		}
	}

	return todoGroups, nil
}

func (d *dal) GetTodoGroupByID(todoGroupID string) (*entities.TodoGroup, error) {
	d.lck.RLocker().Lock()
	defer d.lck.RLocker().Unlock()

	g, ok := d.groups[todoGroupID]
	if !ok {
		return nil, entities.ErrNotFound
	}

	todoGroup := g.toGroupEntity()
	todoGroup.Todos = d.getItemsByGroupID(g.ID)

	return todoGroup, nil
}

func (d *dal) getTodoItemById(todoItemID string) (itemDTO, error) {
	if todoItem, ok := d.items[todoItemID]; ok {
		return todoItem, nil
	}

	return itemDTO{}, entities.ErrNotFound
}

func (d *dal) getItemsByGroupID(groupID string) []*entities.TodoItem {
	var items []*entities.TodoItem

	for _, item := range d.items {
		if item.GroupID == groupID {
			items = append(items, item.toItemEntity())
		}
	}

	return items
}
