package storage

import (
	"context"

	"github.com/hoangphuc3604/todo-list/module/todo/model"
	"github.com/hoangphuc3604/todo-list/util"
)

func (s *TodoStorage) UpdateTodo(ctx context.Context, todo *model.TodoItemUpdate, id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	t, ok := s.todos[id]
	if !ok {
		return util.ErrorNotFound(model.TableNameTodo, id)
	}

	t.Completed = *todo.Completed

	s.todos[id] = t
	return nil
}