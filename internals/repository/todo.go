package repository

import (
	"database/sql"
	"log"
	"sync"
	"time"

	models "myproject/internals/models"
)

type ToDoRepository interface {
	GetAll() ([]*models.ToDo, error)
	GetByID(id string) (*models.ToDo, error)
	Create(todo *models.ToDo) error
	UpdateTitle(id string, title string) error
	UpdateCompleted(id string, completed bool) error
	Delete(id string) error
}

type SQLiteToDoRepository struct {
	db *sql.DB
	mu sync.RWMutex
}

func NewSQLiteToDoRepository(db *sql.DB) *SQLiteToDoRepository {
	return &SQLiteToDoRepository{db: db}
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
		var CreatedAt string
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed, &CreatedAt); err != nil {
			return nil, err
		}
		todo.CreatedAt, _ = time.Parse("2006-01-02 15:04:05.999999999-07:00", CreatedAt)
		todos = append(todos, &todo)
	}
	return todos, nil
}

func (r *SQLiteToDoRepository) GetByID(id string) (*models.ToDo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	row, err := r.db.Query("SELECT * FROM todos WHERE id=?", id)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	todo := models.ToDo{}
	var CreatedAt string
	if err := row.Scan(&todo.ID, &todo.Title, &todo.Completed, &CreatedAt); err != nil {
		return nil, err
	}
	todo.CreatedAt, _ = time.Parse("2006-01-02 15:04:05.999999999-07:00", CreatedAt)
	return &todo, nil
}

func (r *SQLiteToDoRepository) UpdateTitle(id string, title string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := r.db.Exec("UPDATE todos SET title = ? WHERE id = ?", title, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *SQLiteToDoRepository) UpdateCompleted(id string, completed bool) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := r.db.Exec("UPDATE todos SET completed = ? WHERE id = ?", completed, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *SQLiteToDoRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := r.db.Exec("DELETE FROM todos WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
