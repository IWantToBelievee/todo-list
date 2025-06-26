package repository

import (
	"database/sql"
	_ "database/sql"
	_ "fmt"
	"log"
	_ "os"
	"sync"
	_ "sync"

	models "myproject/internals/models"
)

type ToDoRepository interface {
	GetAll() ([]*models.ToDo, error)
	Create(todo *models.ToDo) error
}

type SQLiteToDoRepository struct {
	db *sql.DB
	mu sync.RWMutex
}

func NewSQLiteToDoRepository(db *sql.DB) *SQLiteToDoRepository {
	return &SQLiteToDoRepository{db: db}
}

func (r *SQLiteToDoRepository) GetAll() ([]*models.ToDo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	rows, err := r.db.Query("SELECT * FROM todos")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var todos []*models.ToDo
	for rows.Next() {
		todo := models.ToDo{}
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt); err != nil {
			return nil, err
		}
		todos = append(todos, &todo)
	}
	return todos, nil
}

func (r *SQLiteToDoRepository) Create(todo *models.ToDo) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	result, err := r.db.Exec("INSERT INTO todos (id, title, completed, created_at) VALUES (?, ?, ?, ?)", todo.ID, todo.Title, todo.Completed, todo.CreatedAt)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	todo.ID = string(id)
	return nil
}
