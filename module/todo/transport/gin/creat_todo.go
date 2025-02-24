package gin_todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoangphuc3604/todo-list/module/todo/biz"
	"github.com/hoangphuc3604/todo-list/module/todo/model"
	"github.com/hoangphuc3604/todo-list/module/todo/storage"
	"github.com/hoangphuc3604/todo-list/util"
)

func CreateTask(store *storage.TodoStorage) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var data model.TodoItemCreation

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
			return
		}

		business := biz.NewCreateTodoBiz(store)

		id, err := business.CreateNewTodo(ctx, &data)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusOK, util.CreateTodoResponse(id))
	}
}