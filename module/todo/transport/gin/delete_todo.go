package gin_todo

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hoangphuc3604/todo-list/module/todo/biz"
	"github.com/hoangphuc3604/todo-list/module/todo/storage"
	"github.com/hoangphuc3604/todo-list/util"
)

func DeleteTask(store *storage.TodoStorage) func(*gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, util.ErrorResponse(util.ErrorInvalidRequest(err)))
			return
		}

		business := biz.NewDeleteTodoBiz(store)

		err = business.DeleteTodoById(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusOK, util.UpdateTodoResponse(id))
	}
}