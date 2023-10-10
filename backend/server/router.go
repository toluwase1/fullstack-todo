package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func (s *Server) defineRoutes(router *gin.Engine) {
	router.GET("/ping", s.Ping)
	r := router.Group("/api/v1")

	r.POST("/todos", s.CreateTodo)
	r.GET("/todos/:id", s.GetTodo)
	r.GET("/todos", s.ListTodos)
	r.PUT("/todos/:id", s.UpdateTodo)
	r.DELETE("/todos/:id", s.DeleteTodo)
	r.GET("/todos/search", s.SearchTodoByMultipleParameters)
	//r.GET("/todos/category/:category", s.SearchTodoByCategory)

}

func (s *Server) setupRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	// setup cors
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	s.defineRoutes(r)

	return r
}
