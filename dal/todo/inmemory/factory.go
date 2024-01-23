package inmemory

import (
	"sync"

	"github.com/hobord/poc-htmx-go-todolist/dal/todo"
)

type dal struct {
	lck    sync.RWMutex
	groups map[string]groupDTO
	items  map[string]itemDTO
}

func NewDal() todo.ReaderWriter {
	d := &dal{
		groups: make(map[string]groupDTO),
		items:  make(map[string]itemDTO),
	}

	return d
}
