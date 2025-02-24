package biz

import (
	"context"

	"github.com/hoangphuc3604/todo-list/module/todo/model"
	"github.com/hoangphuc3604/todo-list/util"
)

type DeleteTodoStorage interface {
	DeleteTodo(ctx context.Context, id int) error
}

type deleteTodoBiz struct {
	storage DeleteTodoStorage
}

func NewDeleteTodoBiz(storage DeleteTodoStorage) *deleteTodoBiz {
	return &deleteTodoBiz{storage: storage}
}

func (biz *deleteTodoBiz) DeleteTodoById(ctx context.Context, id int) error {
	if err := biz.storage.DeleteTodo(ctx, id); err != nil {
		return util.ErrorCanNotDeleteEntity(model.TableNameTodo, err)
	}

	return nil
}