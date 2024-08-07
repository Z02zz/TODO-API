package main

import (
	"log"
	_ "net/http"

	"github.com/gin-gonic/gin"
	_ "gorm.io/gorm"

	"TODO-API/internal/handler"
	"TODO-API/internal/repository"
	"TODO-API/pkg/database"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	todoRepo := repository.NewTodoRepository(db)
	todoHandler := handler.NewTodoHandler(todoRepo)

	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/todos", todoHandler.CreateTodo)
	r.GET("/todos", todoHandler.GetAllTodos)
	r.GET("/todos/:id", todoHandler.GetTodo)
	r.PUT("/todos/:id", todoHandler.UpdateTodo)
	r.DELETE("/todos/:id", todoHandler.DeleteTodo)

	log.Fatal(r.Run(":8080"))
}
