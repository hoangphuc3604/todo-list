package biz

import (
	"context"

	"github.com/hoangphuc3604/todo-list/module/todo/model"
	"github.com/hoangphuc3604/todo-list/util"
)

type UpdateTodoStorage interface {
	UpdateTodo(ctx context.Context, todo *model.TodoItemUpdate, id int) error
}

type updateTodoBiz struct {
	storage UpdateTodoStorage
}

func NewUpdateTodoBiz(storage UpdateTodoStorage) *updateTodoBiz {
	return &updateTodoBiz{storage: storage}
}

func (biz *updateTodoBiz) UpdateTodoByID(ctx context.Context, todo *model.TodoItemUpdate, id int) error {
	err := todo.Validate()
	if err != nil {
		return util.ErrorInvalidRequest(err)
	}

	return biz.storage.UpdateTodo(ctx, todo, id)
}