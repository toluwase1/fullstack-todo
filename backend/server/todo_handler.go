package server

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"todo/models"
)

func (s *Server) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "we are live",
	})
}
func (s *Server) CreateTodo(c *gin.Context) {
	var input models.TodoRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := validateTodoInput(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := models.Todo{
		Title:       input.Title,
		Description: input.Description,
		Category:    input.Category,
		Status:      input.Status,
	}
	createdTodo, err := s.TodoService.CreateTodo(todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Todo"})
		return
	}

	c.JSON(http.StatusCreated, createdTodo)
}

func (s *Server) GetTodo(c *gin.Context) {
	todoID := c.Param("id")

	todo, err := s.TodoService.FindTodoById(todoID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (s *Server) ListTodos(c *gin.Context) {
	todos, err := s.TodoService.FindAllTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve todos"})
		return
	}
	c.JSON(http.StatusOK, todos)
}

func (s *Server) UpdateTodo(c *gin.Context) {
	todoID := c.Param("id")
	var updateRequest models.TodoRequest

	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTodo, err := s.TodoService.UpdateTodoByID(todoID, updateRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update todo"})
		return
	}

	c.JSON(http.StatusOK, updatedTodo)
}

func (s *Server) DeleteTodo(c *gin.Context) {
	todoID := c.Param("id")

	err := s.TodoService.DeleteTodoByID(todoID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found", "err": err})
		return
	}
	result := "item with ID " + todoID + " deleted"
	c.JSON(http.StatusOK, result)
}

func (s *Server) SearchTodoByMultipleParameters(c *gin.Context) {
	title := c.DefaultQuery("title", "")
	category := c.DefaultQuery("category", "")
	description := c.DefaultQuery("description", "")
	status := c.DefaultQuery("status", "")

	searchRequest := models.TodoRequest{
		Title:       title,
		Description: description,
		Category:    category,
		Status:      status,
	}

	todos, err := s.TodoService.SearchTodosByMultipleParameters(searchRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search for todos"})
		return
	}

	c.JSON(http.StatusOK, todos)
}

//func (s *Server) SearchTodoByCategory(c *gin.Context) {
//	categoryName := c.Param("category")
//
//	todos, err := s.TodoService.SearchTodosByCategory(categoryName)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search for todos"})
//		return
//	}
//
//	c.JSON(http.StatusOK, todos)
//}

func validateTodoInput(input models.TodoRequest) error {
	if strings.TrimSpace(input.Title) == "" {
		return errors.New("title is required")
	}
	if strings.TrimSpace(input.Description) == "" {
		return errors.New("description is required")
	}
	if strings.TrimSpace(input.Category) == "" {
		return errors.New("category is required")
	}
	if strings.TrimSpace(input.Status) == "" {
		return errors.New("status is required")
	}

	//status validation
	if input.Status != models.InProgress &&
		input.Status != models.StatusPending &&
		input.Status != models.StatusDone {
		return errors.New("invalid status")
	}

	//category validation
	if input.Category != models.CategoryPersonal &&
		input.Category != models.CategoryWork &&
		input.Category != models.CategoryHome {
		return errors.New("invalid category")
	}
	return nil
}
