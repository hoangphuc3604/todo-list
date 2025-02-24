package biz

import (
	"context"
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
	return biz.storage.DeleteTodo(ctx, id)
}