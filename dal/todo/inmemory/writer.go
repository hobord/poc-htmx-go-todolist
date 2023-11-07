package inmemory

import (
	"github.com/google/uuid"
	"github.com/hobord/poc-htmx-go-todolist/entities"
)

func generateID() string {
	id := uuid.New()
	return id.String()
}

func (d *dal) Create(todo *entities.Todo) error {
	d.lck.Lock()
	defer d.lck.Unlock()

	if d.data[todo.User] == nil {
		d.data[todo.User] = make(map[string][]*entities.Todo)
	}

	if d.data[todo.User][todo.Group] == nil {
		d.data[todo.User][todo.Group] = make([]*entities.Todo, 0, 10)
	}

	todo.ID = generateID()

	d.data[todo.User][todo.Group] = append(d.data[todo.User][todo.Group], todo)

	return nil
}

func (d *dal) Update(todo *entities.Todo) error {
	d.lck.Lock()
	defer d.lck.Unlock()

	if d.data[todo.User] == nil || d.data[todo.User][todo.Group] == nil {
		return nil
	}

	originalTodo, err := d.GetByID(todo.ID)
	if err != nil {
		return err
	}

	if originalTodo.Group != todo.Group {
		d.data[originalTodo.User][originalTodo.Group] = remove(d.data[originalTodo.User][originalTodo.Group], originalTodo)
	}

	d.data[todo.User][todo.Group] = append(d.data[todo.User][todo.Group], todo)

	return nil
}

func remove(todos []*entities.Todo, todo *entities.Todo) []*entities.Todo {
	for i, t := range todos {
		if t.ID == todo.ID {
			return append(todos[:i], todos[i+1:]...)
		}
	}
	return todos
}

func (d *dal) Delete(id string) error {
	d.lck.Lock()
	defer d.lck.Unlock()

	tudo, err := d.GetByID(id)
	if err != nil {
		return err
	}

	d.data[tudo.User][tudo.Group] = remove(d.data[tudo.User][tudo.Group], tudo)

	return nil
}

func (d *dal) SetCompleted(id string, completed bool) error {
	d.lck.Lock()
	defer d.lck.Unlock()

	todo, err := d.GetByID(id)
	if err != nil {
		return err
	}

	todo.Completed = completed

	return nil
}

func (d *dal) SetPriority(todos []*entities.Todo) error {
	d.lck.Lock()
	defer d.lck.Unlock()

	// TODO: implement

	return nil
}
