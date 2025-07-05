package repository

import (
	models "myproject/internals/models"
)

type ToDoRepository interface {
	Create(todo *models.ToDo) error
	GetAll() ([]*models.ToDo, error)
	GetByID(id *string) (*models.ToDo, error)
	Update(req *models.UpdateRequest) error
	Delete(id *string) error
	SaveOrder(order []*string) error
	GetOrder() ([]*string, error)
}
