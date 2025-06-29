package db

import (
	"database/sql"

	_ "github.com/glebarez/go-sqlite"
)

func InitDB(filepath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", filepath)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	// Инициализация схемы
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS todos (
            id TEXT PRIMARY KEY,
            title TEXT NOT NULL,
            completed INTEGER NOT NULL,
			created_at TEXT NOT NULL,
			order_id INTEGER AUTOINCREMENT
        )
    `)
	if err != nil {
		return nil, err
	}

	return db, nil
}
