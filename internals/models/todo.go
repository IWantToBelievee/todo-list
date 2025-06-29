package models

import "time"

type ToDo struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	OrderId   int       `json:"order_id"`
}
