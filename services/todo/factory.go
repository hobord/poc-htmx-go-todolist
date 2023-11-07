package todo

import "github.com/hobord/poc-htmx-go-todolist/dal/todo"

type service struct {
	dal todo.ReaderWriter
}

func NewService(dal todo.ReaderWriter) (Service, error) {
	s := &service{
		dal: dal,
	}

	return s, s.validate()
}

func (s *service) validate() error {
	if s.dal == nil {
		return ErrDalIsNil
	}

	return nil
}
