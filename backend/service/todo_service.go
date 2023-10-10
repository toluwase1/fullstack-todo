package service

import (
	"errors"
	"fmt"
	"todo/db"
	"todo/models"
)

// TodoService interface
type TodoService interface {
	CreateTodo(request models.Todo) (models.Todo, error)
	FindTodoById(id string) (models.Todo, error)
	DeleteTodoByID(id string) error
	FindAllTodos() ([]models.Todo, error)
	UpdateTodoByID(todoId string, updateRequest models.TodoRequest) (models.Todo, error)
	SearchTodosByMultipleParameters(searchRequest models.TodoRequest) ([]models.Todo, error)
	SearchTodosByCategory(category string) ([]models.Todo, error)
}

// todoService struct
type todoService struct {
	todoRepo db.TodoRepository
}

func NewTodoService(todoRepo db.TodoRepository) TodoService {
	return &todoService{
		todoRepo: todoRepo,
	}
}

func (t todoService) CreateTodo(request models.Todo) (models.Todo, error) {
	todo, err := t.todoRepo.CreateTodo(request)
	if err != nil {
		return models.Todo{}, err
	}
	return todo, nil
}

func (t todoService) FindTodoById(id string) (models.Todo, error) {
	todo, err := t.todoRepo.FindTodoByID(id)
	if err != nil {
		return models.Todo{}, err
	}
	return todo, nil
}

func (t todoService) DeleteTodoByID(id string) error {

	_, err := t.todoRepo.FindTodoByID(id)
	if err != nil {
		msg := fmt.Sprintf("item with id " + id + "not found")
		return errors.New(msg)
	}

	err = t.todoRepo.DeleteTodoByID(id)
	if err != nil {
		return err
	}
	return nil
}

func (t todoService) FindAllTodos() ([]models.Todo, error) {
	todos, err := t.todoRepo.GetAllTodos()
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (t todoService) UpdateTodoByID(todoId string, request models.TodoRequest) (models.Todo, error) {
	todo, err := t.todoRepo.UpdateTodoByID(todoId, request)
	if err != nil {
		return models.Todo{}, err
	}
	return todo, nil
}

func (t todoService) SearchTodosByMultipleParameters(searchRequest models.TodoRequest) ([]models.Todo, error) {
	todos, err := t.todoRepo.SearchTodosByMultipleParameters(searchRequest)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (t todoService) SearchTodosByCategory(category string) ([]models.Todo, error) {
	categories, err := t.todoRepo.SearchTodosByCategory(category)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
