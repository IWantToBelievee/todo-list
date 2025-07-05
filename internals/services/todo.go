package services

import (
	"myproject/internals/models"
	"myproject/internals/repository"
	"time"

	"github.com/google/uuid"
)

type ToDoService struct {
	repo repository.ToDoRepository
}

func NewToDoService() (*ToDoService, error) {
	repo, err := repository.NewJsonToDoRepository("todos-data.json")
	if err != nil {
		return nil, err
	}
	return &ToDoService{repo: repo}, nil
}

func (s *ToDoService) Create() error {
	todo := &models.ToDo{
		ID:        uuid.NewString(),
		Title:     "Undefined",
		Completed: false,
		CreatedAt: time.Now(),
	}
	if err := s.repo.Create(todo); err != nil {
		return err
	}
	return nil
}

func (s *ToDoService) GetAll() ([]*models.ToDo, error) {
	return s.repo.GetAll()
}

func (s *ToDoService) GetByID(id *string) (*models.ToDo, error) {
	return s.repo.GetByID(id)
}

func (s *ToDoService) Update(req *models.UpdateRequest) error {
	if req != nil {
		if req.Title != nil || req.Description != nil || req.Completed != nil {
			return s.repo.Update(req)
		}
	}
	return nil
}

func (s *ToDoService) Delete(id *string) error {
	return s.repo.Delete(id)
}

func (s *ToDoService) SaveOrder(order []*string) error {
	return s.repo.SaveOrder(order)
}

func (s *ToDoService) GetOrder() ([]*string, error) {
	return s.repo.GetOrder()
}
