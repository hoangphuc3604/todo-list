package biz

import (
	"context"

	"github.com/hoangphuc3604/todo-list/module/todo/model"
)

type ListAllTodoStorage interface {
	ListTodo(ctx context.Context) ([]model.TodoItem, error)
}

type listAllTodoBiz struct {
	storage ListAllTodoStorage
}

func NewListAllTodoBiz(storage ListAllTodoStorage) *listAllTodoBiz {
	return &listAllTodoBiz{storage: storage}
}

func (biz *listAllTodoBiz) ListAllTodo(ctx context.Context) ([]model.TodoItem, error) {
	return biz.storage.ListTodo(ctx)
}