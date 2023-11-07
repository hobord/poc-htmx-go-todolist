package todo

import "github.com/hobord/poc-htmx-go-todolist/entities"

func (s *service) Create(user, group, title string) (*entities.Todo, error) {
	todo := &entities.Todo{
		User:  user,
		Group: group,
		Title: title,
	}

	if err := s.dal.Create(todo); err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *service) GetAll(user string) ([]*entities.Todo, error) {
	return s.dal.GetAll(user)
}

func (s *service) GetByGroup(user string, group string) ([]*entities.Todo, error) {
	return s.dal.GetByGroup(user, group)
}

func (s *service) GetAllGroup(user string) (map[string][]*entities.Todo, error) {
	groups := make(map[string][]*entities.Todo)

	todos, err := s.dal.GetAll(user)
	if err != nil {
		return nil, err
	}

	for _, todo := range todos {
		if groups[todo.Group] == nil {
			groups[todo.Group] = make([]*entities.Todo, 0, 10)
		}

		groups[todo.Group] = append(groups[todo.Group], todo)
	}

	return groups, nil
}
