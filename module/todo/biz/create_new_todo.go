package biz

import (
	"context"

	"github.com/hoangphuc3604/todo-list/module/todo/model"
	"github.com/hoangphuc3604/todo-list/util"
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
	err := todo.Validate()
	if err != nil {
		return 0, util.ErrorInvalidRequest(err)
	}

	id, err := biz.storage.CreateTodo(ctx, todo)
	if err != nil {
		return id, util.ErrorCanNotCreateEntity(model.TableNameTodo, err)
	}

	return id, nil
}