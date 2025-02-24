package gin_todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoangphuc3604/todo-list/module/todo/biz"
	"github.com/hoangphuc3604/todo-list/module/todo/storage"
	"github.com/hoangphuc3604/todo-list/util"
)

func ListTodo(store *storage.TodoStorage) func(*gin.Context) {
	return func(ctx *gin.Context) {
		business := biz.ListAllTodoStorage(store)

		data, err := business.ListTodo(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusOK, data)
	}
}