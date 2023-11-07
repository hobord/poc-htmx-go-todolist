package composition

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/hobord/poc-htmx-go-todolist/entities"
	"github.com/hobord/poc-htmx-go-todolist/services/health"
	"github.com/hobord/poc-htmx-go-todolist/services/todo"
	"github.com/stretchr/testify/mock"
)

func createMockTodoService() todo.Service {
	service := &todo.MockService{}

	todoItems := make([]*entities.Todo, 0, 10)
	for i := 1; i <= 10; i++ {
		todoItems = append(todoItems, &entities.Todo{
			ID:          fmt.Sprintf("%d", i),
			Title:       fmt.Sprintf("Title %d", i),
			Description: fmt.Sprintf("Description %d", i),
			Group:       "Group1",
			Completed:   randomBoolGenerator(),
			Priority:    1,
		})
	}

	service.On("GetAllGroup", mock.Anything).Return(
		map[string][]*entities.Todo{
			"Group1": todoItems,
		}, nil)

	service.On("GetByGroup", mock.Anything, mock.Anything).Return(
		todoItems, nil)

	service.On("Create", mock.Anything).Return(nil)

	return service
}

func randomBoolGenerator() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2) == 1
}

func createMockHealtService() health.Service {
	service := &health.MockService{}

	service.On("Health").Return(nil)

	return service
}
