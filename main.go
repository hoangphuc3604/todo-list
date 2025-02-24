package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hoangphuc3604/todo-list/module/todo/storage"
	gin_todo "github.com/hoangphuc3604/todo-list/module/todo/transport/gin"
)

func main() {
	r := gin.Default()
	store := storage.NewTodoStorage()

	task := r.Group("/tasks")
	{
		task.POST("", gin_todo.CreateTask(store))
		task.GET("", gin_todo.ListTodo(store))
	}

	r.Run(":8080")
}