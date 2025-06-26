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
		return &App{todoSvc: nil}
	}
	return &App{todoSvc: todoSvc}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) shutdown(ctx context.Context) {}

func (a *App) GetToDos() ([]*models.ToDo, error) {
	return a.todoSvc.GetAll()
}

func (a *App) CreateToDo(title string) (*models.ToDo, error) {
	return a.todoSvc.Create(title)
}
