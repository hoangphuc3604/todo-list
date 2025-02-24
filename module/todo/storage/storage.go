package storage

import (
	"sync"

	"github.com/hoangphuc3604/todo-list/module/todo/model"
)

type TodoStorage struct {
	mu sync.Mutex
	todos map[int]model.TodoItem
	nextID int
}

func NewTodoStorage() *TodoStorage {
	return &TodoStorage{
		todos: make(map[int]model.TodoItem),
		nextID: 1,
	}
}

