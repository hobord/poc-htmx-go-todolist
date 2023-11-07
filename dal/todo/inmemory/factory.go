package inmemory

import (
	"sync"

	"github.com/hobord/poc-htmx-go-todolist/dal/todo"
	"github.com/hobord/poc-htmx-go-todolist/entities"
)

type dal struct {
	lck  sync.RWMutex
	data map[string]map[string][]*entities.Todo
}

func NewDal() todo.ReaderWriter {
	d := &dal{
		data: make(map[string]map[string][]*entities.Todo),
	}

	d.data["user1"] = make(map[string][]*entities.Todo)
	d.data["user1"]["ds"] = []*entities.Todo{
		{
			ID:    "1",
			Title: "ds 1",
			Group: "ds",
		},
	}

	return d
}
