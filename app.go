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
		panic("failed to initialize ToDoService: " + err.Error())
	}
	return &App{todoSvc: todoSvc}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) shutdown(ctx context.Context) {}

func (a *App) CreateToDo() error {
	return a.todoSvc.Create()
}

func (a *App) GetToDos() ([]*models.ToDo, error) {
	return a.todoSvc.GetAll()
}

func (a *App) GetToDoByID(id *string) (*models.ToDo, error) {
	return a.todoSvc.GetByID(id)
}

func (a *App) UpdateToDo(req *models.UpdateRequest) error {
	return a.todoSvc.Update(req)
}

func (a *App) DeleteToDo(id *string) error {
	return a.todoSvc.Delete(id)
}

func (a *App) SaveOrder(order []*string) error {
	return a.todoSvc.SaveOrder(order)
}

func (a *App) GetOrder() ([]*string, error) {
	return a.todoSvc.GetOrder()
}
