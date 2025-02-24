package storage

import (
	"context"

	"github.com/hoangphuc3604/todo-list/module/todo/model"
)

func (s *TodoStorage) ListTodo(ctx context.Context) ([]model.TodoItem, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	todos := make([]model.TodoItem, 0, len(s.todos))
	for _, todo := range s.todos {
		todos = append(todos, todo)
	}
	return todos, nil
}