package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {

	var err error

	DB, err = sql.Open("sqlite3", "database/todos.db")
	if err != nil {
		panic(err.Error())
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()

}

func createTables() {

	createTodosTable := `
			CREATE TABLE IF NOT EXISTS todos (
    			id INTEGER PRIMARY KEY AUTOINCREMENT,
    			title TEXT NOT NULL,
    			description TEXT,
    			status BOOLEAN DEFAULT FALSE,
    			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
			)	
	`
	_, err := DB.Exec(createTodosTable)

	if err != nil {
		panic(err.Error())
	}
}
