package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hoangphuc3604/todo-list/module/todo/storage"
	gin_todo "github.com/hoangphuc3604/todo-list/module/todo/transport/gin"
)

func main() {
	r := gin.Default()
	store := storage.NewTodoStorage()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000", "https://todo-list-ui-c5e6.onrender.com"},
		AllowMethods: []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
	}))

	task := r.Group("/tasks")
	{
		task.POST("", gin_todo.CreateTask(store))
		task.GET("", gin_todo.ListTodo(store))
		task.PUT("/:id", gin_todo.UpdateTask(store))
		task.DELETE("/:id", gin_todo.DeleteTask(store))
	}

	r.Run(":8080")
}