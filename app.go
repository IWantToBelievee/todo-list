package main

import (
	"context"
	"myproject/internals/models"
	"myproject/internals/services"
)

type App struct {
	ctx     context.Context
	todoSvc *services.ToDoService
}

func NewApp() *App {
	todoSvc, err := services.NewToDoService()
	if err != nil {
		panic("failed to initialize ToDoService: " + err.Error()) // Or use a logger
	}
	return &App{todoSvc: todoSvc}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) shutdown(ctx context.Context) {}

func (a *App) CreateToDo(title string) (*models.ToDo, error) {
	return a.todoSvc.Create(title)
}

func (a *App) GetToDos() ([]*models.ToDo, error) {
	return a.todoSvc.GetAll()
}

func (a *App) GetToDoByID(id string) (*models.ToDo, error) {
	return a.todoSvc.GetByID(id)
}

func (a *App) UpdateToDoTitle(id string, title string) error {
	return a.todoSvc.UpdateTitle(id, title)
}

func (a *App) UpdateToDoCompleted(id string, completed bool) error {
	return a.todoSvc.UpdateCompleted(id, completed)
}

func (a *App) DeleteToDo(id string) error {
	return a.todoSvc.Delete(id)
}
