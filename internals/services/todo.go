package services

import (
	"errors"
	"myproject/internals/db"
	"myproject/internals/models"
	"myproject/internals/repository"
	"time"

	"github.com/google/uuid"
)

type ToDoService struct {
	repo repository.ToDoRepository
}

func NewToDoService() (*ToDoService, error) {
	db_, err := db.InitDB("./todos.db")
	if err != nil {
		return nil, err
	}
	repo := repository.NewSQLiteToDoRepository(db_)

	return &ToDoService{repo: repo}, nil
}

func (s *ToDoService) Create(title string) (*models.ToDo, error) {
	if title == "" {
		return nil, errors.New("title cannot be empty")
	}
	todo := &models.ToDo{
		ID:        uuid.NewString(),
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	}
	if err := s.repo.Create(todo); err != nil {
		return nil, err
	}
	return todo, nil
}

func (s *ToDoService) GetAll() ([]*models.ToDo, error) {
	return s.repo.GetAll()
}

func (s *ToDoService) GetByID(id string) (*models.ToDo, error) {
	return s.repo.GetByID(id)
}

func (s *ToDoService) UpdateTitle(id string, title string) error {
	return s.repo.UpdateTitle(id, title)
}

func (s *ToDoService) UpdateCompleted(id string, completed bool) error {
	return s.repo.UpdateCompleted(id, completed)
}

func (s *ToDoService) Delete(id string) error {
	return s.repo.Delete(id)
}
