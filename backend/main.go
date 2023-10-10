package main

import (
	"net/http"
	"time"
	"todo/db"
	"todo/server"
	"todo/service"
)

func main() {
	http.DefaultClient.Timeout = time.Second * 10
	gormDB := db.GetDB()
	todoRepo := db.NewTodoRepository(gormDB)
	todoService := service.NewTodoService(todoRepo)

	s := &server.Server{
		TodoRepository: todoRepo,
		TodoService:    todoService,
	}
	s.Start()
}
