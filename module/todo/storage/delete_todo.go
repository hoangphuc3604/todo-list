package storage

import (
	"context"

	"github.com/hoangphuc3604/todo-list/module/todo/model"
	"github.com/hoangphuc3604/todo-list/util"
)

func (s *TodoStorage) DeleteTodo(ctx context.Context, id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.todos[id]; !ok {
		return util.ErrorNotFound(model.TableNameTodo, id)
	}

	delete(s.todos, id)
	return nil
}