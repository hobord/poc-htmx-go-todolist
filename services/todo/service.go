package todo

import (
	"github.com/hobord/poc-htmx-go-todolist/entities"
)

func (s *service) GetTodoGroupsByUserID(userID string) ([]*entities.TodoGroup, error) {
	return s.dal.GetTodoGroupsByUserID(userID)
}

func (s *service) GetTodoGroupByID(todoGroupID string) (*entities.TodoGroup, error) {
	return s.dal.GetTodoGroupByID(todoGroupID)
}

func (s *service) CreateTodoGroup(todoGroup *entities.TodoGroup) error {
	return s.dal.WriteTodoGroup(todoGroup)
}

func (s *service) UpdateTodoGroup(todoGroup *entities.TodoGroup) error {
	return s.dal.WriteTodoGroup(todoGroup)
}

func (s *service) DeleteTodoGroup(groupID string) error {
	return s.dal.DeleteTodoGroup(groupID)
}

func (s *service) AddTodoItem(todoItem *entities.TodoItem) error {
	return s.dal.WriteTodoItem(todoItem)
}

func (s *service) UpdateTodoItem(todoItem *entities.TodoItem) error {
	return s.dal.WriteTodoItem(todoItem)
}

func (s *service) SortTodoItems(ids []string) error {
	for i, id := range ids {
		item, err := s.dal.GetTodoItemByID(id)
		if err != nil {
			return err
		}

		item.Priority = i

		if err := s.dal.WriteTodoItem(item); err != nil {
			return err
		}
	}

	return nil
}

func (s *service) DeleteTodoItem(todoItemID string) error {
	return s.dal.DeleteTodoItem(todoItemID)
}

func (s *service) DeleteCompletedTodoItems(groupID string) error {
	group, err := s.dal.GetTodoGroupByID(groupID)
	if err != nil {
		return err
	}

	for _, todo := range group.Todos {
		if todo.Completed {
			if err := s.dal.DeleteTodoItem(todo.ID); err != nil {
				return err
			}
		}
	}

	return nil
}
