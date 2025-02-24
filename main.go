package main

import (
	"github.com/gin-gonic/gin"
	gin_todo "github.com/hoangphuc3604/todo-list/module/todo/transport/gin"
)

func main() {
	r := gin.Default()

	task := r.Group("/tasks")
	{
		task.POST("", gin_todo.CreateTask())
	}

	r.Run(":8080")
}