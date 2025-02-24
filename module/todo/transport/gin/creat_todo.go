package gin_todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoangphuc3604/todo-list/module/todo/biz"
	"github.com/hoangphuc3604/todo-list/module/todo/model"
	"github.com/hoangphuc3604/todo-list/module/todo/storage"
)

func CreateTask(store *storage.TodoStorage) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var data model.TodoItemCreation

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		business := biz.NewCreateTodoBiz(store)

		id, err := business.CreateNewTodo(ctx, &data)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"id": id,
		})
	}
}