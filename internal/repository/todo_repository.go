package repository

import (
	"TODO-API/internal/model"

	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

func (r *TodoRepository) Create(todo *model.Todo) error {
	return r.db.Create(todo).Error
}

func (r *TodoRepository) GetAll() ([]model.Todo, error) {
	var todos []model.Todo
	err := r.db.Find(&todos).Error
	return todos, err
}

func (r *TodoRepository) GetByID(id uint) (*model.Todo, error) {
	var todo model.Todo
	err := r.db.First(&todo, id).Error
	return &todo, err
}

func (r *TodoRepository) Update(todo *model.Todo) error {
	return r.db.Save(todo).Error
}

func (r *TodoRepository) Delete(id uint) error {
	return r.db.Delete(&model.Todo{}, id).Error
}
