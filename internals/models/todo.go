package models

import "time"

type ToDo struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

func (t *ToDo) SetTitle(title *string) {
	t.Title = *title
}

func (t *ToDo) SetComplited(completed *bool) {
	t.Completed = *completed
}
