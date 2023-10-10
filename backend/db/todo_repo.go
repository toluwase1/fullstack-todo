package db

import (
	"gorm.io/gorm"
	"todo/models"
)

type todoRepository struct {
	DB *gorm.DB
}

func NewTodoRepository(db *GormDB) TodoRepository {
	return &todoRepository{db.DB}
}

type TodoRepository interface {
	CreateTodo(todo models.Todo) (models.Todo, error)
	FindTodoByID(id string) (models.Todo, error)
	DeleteTodoByID(id string) error
	UpdateTodoByID(ID string, request models.TodoRequest) (models.Todo, error)
	GetAllTodos() ([]models.Todo, error)
	SearchTodosByMultipleParameters(searchRequest models.TodoRequest) ([]models.Todo, error)
	SearchTodosByCategory(category string) ([]models.Todo, error)
}

func (t *todoRepository) CreateTodo(todo models.Todo) (models.Todo, error) {
	if err := t.DB.Create(&todo).Error; err != nil {
		return models.Todo{}, err
	}
	return todo, nil
}

func (t *todoRepository) FindTodoByID(id string) (models.Todo, error) {
	var todo models.Todo
	if err := t.DB.Where("id = ?", id).First(&todo).Error; err != nil {
		return models.Todo{}, err
	}
	return todo, nil
}

func (t *todoRepository) DeleteTodoByID(id string) error {
	if err := t.DB.Where("id = ?", id).Delete(&models.Todo{}).Error; err != nil {
		return err
	}
	return nil
}

func (t *todoRepository) UpdateTodoByID(ID string, request models.TodoRequest) (models.Todo, error) {
	var todo models.Todo
	if err := t.DB.Where("id = ?", ID).First(&todo).Error; err != nil {
		return models.Todo{}, err
	}

	t.DB.Model(&todo).Updates(request)

	return todo, nil
}

func (t *todoRepository) GetAllTodos() ([]models.Todo, error) {
	var todos []models.Todo
	if err := t.DB.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (t *todoRepository) SearchTodosByMultipleParameters(searchRequest models.TodoRequest) ([]models.Todo, error) {
	var todos []models.Todo

	// Start with a fresh query builder for each search
	query := t.DB

	if searchRequest.Title != "" {
		query = query.Where("title ILIKE ?", "%"+searchRequest.Title+"%")
	}
	if searchRequest.Description != "" {
		query = query.Where("description ILIKE ?", "%"+searchRequest.Description+"%")
	}
	if searchRequest.Category != "" {
		query = query.Where("category ILIKE ?", "%"+searchRequest.Category+"%")
	}
	if searchRequest.Status != "" {
		query = query.Where("status = ?", searchRequest.Status)
	}

	// Execute the query and check for errors
	if err := query.Find(&todos).Error; err != nil {
		return nil, err
	}

	return todos, nil
}

func (t *todoRepository) SearchTodosByCategory(category string) ([]models.Todo, error) {
	var todos []models.Todo
	if err := t.DB.Where("category = ?", category).Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}
