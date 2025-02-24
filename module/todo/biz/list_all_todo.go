package biz

import (
	"context"

	"github.com/hoangphuc3604/todo-list/module/todo/model"
	"github.com/hoangphuc3604/todo-list/util"
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
	todos, err := biz.storage.ListTodo(ctx)
	if err != nil {
		return nil, util.ErrorCanNotListEntity(model.TableNameTodo, err)
	}
	return todos, nil
}