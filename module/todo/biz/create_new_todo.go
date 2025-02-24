package biz

import (
	"context"

	"github.com/hoangphuc3604/todo-list/module/todo/model"
)

type CreateTodoStorage interface {
	CreateTodo(ctx context.Context, todo *model.TodoItemCreation) (int, error)
}

type createTodoBiz struct {
	storage CreateTodoStorage
}

func NewCreateTodoBiz(storage CreateTodoStorage) *createTodoBiz {
	return &createTodoBiz{storage: storage}
}

func (biz *createTodoBiz) CreateNewTodo(ctx context.Context, todo *model.TodoItemCreation) (int, error) {
	return biz.storage.CreateTodo(ctx, todo)
}