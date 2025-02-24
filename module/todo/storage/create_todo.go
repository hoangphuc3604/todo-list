package storage

import (
	"context"

	"github.com/hoangphuc3604/todo-list/module/todo/model"
)

func (s *TodoStorage) CreateTodo(ctx context.Context, todo *model.TodoItemCreation) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	newTodo := model.TodoItem{
		ID:          s.nextID,
		Title:       todo.Title,
		Description: todo.Description,
		Completed:   false,
	}

	s.todos[s.nextID] = newTodo
	s.nextID++
	return newTodo.ID, nil
}