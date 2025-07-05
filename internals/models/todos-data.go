package models

import (
	"errors"
)

type TodosData struct {
	Todos []*ToDo
	Order []*string
}

func OrderContains(id *string, order []*string) bool {
	for _, item := range order {
		if item == id {
			return true
		}
	}
	return false
}

func TodosContains(id *string, todos []*ToDo) bool {
	for _, todo := range todos {
		if todo.ID == *id {
			return true
		}
	}
	return false
}

func (td *TodosData) FixOrder() {
	if len(td.Todos) > len(td.Order) {
		for _, todo := range td.Todos {
			if !OrderContains(&todo.ID, td.Order) {
				td.Order = append(td.Order, &todo.ID)
			}
		}
	} else if len(td.Todos) < len(td.Order) {
		for i, id := range td.Order {
			if !TodosContains(id, td.Todos) {
				td.Order = append(td.Order[:i], td.Order[i+1:]...)
			}
		}
	}
}

func (td *TodosData) SetOrder(order []*string) {
	td.Order = order
	td.FixOrder()
}

func (td *TodosData) AddTodo(todo *ToDo) error {
	if todo != nil {
		td.Todos = append(td.Todos, todo)
		td.Order = append(td.Order, &todo.ID)
		return nil
	}
	return errors.New("todo cannot be nil")
}

func (td *TodosData) GetTodos() []*ToDo {
	return td.Todos
}

func (td *TodosData) GetOrder() []*string {
	return td.Order
}

func (td *TodosData) GetByID(id *string) (*ToDo, *int, error) {
	for i, item := range td.Todos {
		if item.ID == *id {
			return item, &i, nil
		}
	}
	return nil, nil, errors.New("there is no ToDo with id \"" + *id + "\"")
}

func (td *TodosData) UpdateTodoByReq(req *UpdateRequest) error {
	_, idx, err := td.GetByID(&req.ID)
	if err != nil {
		return err
	}

	if req.Title != nil {
		td.Todos[*idx].SetTitle(req.Title)
	}
	if req.Completed != nil {
		td.Todos[*idx].SetComplited(req.Completed)
	}

	return nil
}

func (td *TodosData) RemoveTodo(id *string) error {
	for i, item := range td.Todos {
		if item.ID == *id {
			td.Todos = append(td.Todos[:i], td.Todos[i+1:]...)
			td.FixOrder()
			return nil
		}
	}
	return errors.New("there is no ToDo with id \"" + *id + "\"")
}
