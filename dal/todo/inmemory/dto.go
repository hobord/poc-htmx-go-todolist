package inmemory

import "github.com/hobord/poc-htmx-go-todolist/entities"

type groupDTO struct {
	ID     string
	UserID string
	Title  string
	Color  string
}

func toGroupDTO(todoGroup *entities.TodoGroup) groupDTO {
	return groupDTO{
		ID:     todoGroup.ID,
		UserID: todoGroup.UserID,
		Title:  todoGroup.Title,
		Color:  todoGroup.Color,
	}
}

func (g groupDTO) toGroupEntity() *entities.TodoGroup {
	return &entities.TodoGroup{
		ID:     g.ID,
		UserID: g.UserID,
		Title:  g.Title,
		Color:  g.Color,
	}
}

type itemDTO struct {
	ID        string
	GroupID   string
	Title     string
	Completed bool
	Priority  int
}

func (i itemDTO) toItemEntity() *entities.TodoItem {
	return &entities.TodoItem{
		ID:        i.ID,
		GroupID:   i.GroupID,
		Title:     i.Title,
		Completed: i.Completed,
		Priority:  i.Priority,
	}
}

func toItemDTO(todoItem *entities.TodoItem) itemDTO {
	return itemDTO{
		ID:        todoItem.ID,
		GroupID:   todoItem.GroupID,
		Title:     todoItem.Title,
		Completed: todoItem.Completed,
		Priority:  todoItem.Priority,
	}
}
