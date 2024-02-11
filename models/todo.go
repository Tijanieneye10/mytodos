package models

import (
	"errors"
	"github.com/Tijanieneye10/database"
	"time"
)

type Todo struct {
	ID          int64
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Status      bool      `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

func (t *Todo) Save() (*Todo, error) {

	query := `INSERT INTO todos(title, description, status, created_at) VALUES(?,?,?,?)`

	stmt, err := database.DB.Prepare(query)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	defer stmt.Close()

	t.CreatedAt = time.Now()
	t.Status = false

	result, err := stmt.Exec(&t.Title, &t.Description, &t.Status, &t.CreatedAt)
	if err != nil {
		return nil, err
	}

	t.ID, _ = result.LastInsertId()

	return t, nil
}

func GetTodos() ([]Todo, error) {

	query := `SELECT * FROM todos`

	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var todos []Todo

	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Status, &todo.CreatedAt)
		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	return todos, nil

}

func GetTodo(todoId int64) (*Todo, error) {
	query := `SELECT * FROM todos WHERE id = ?`
	row, err := database.DB.Query(query, todoId)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	var todo *Todo

	err = row.Scan(&todo.ID, &todo.Title, &todo.Status, &todo.CreatedAt)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return todo, nil
}

func MarkTodoCompleted(todoId int64) error {
	query := `UPDATE todos SET status = true WHERE id = ?`

	stmt, err := database.DB.Prepare(query)

	if err != nil {
		return errors.New(err.Error())
	}

	_, err = stmt.Exec(todoId)
	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func UndoTodo(todoId int64) error {

	query := `UPDATE todos SET status = false WHERE id = ?`

	stmt, err := database.DB.Prepare(query)

	if err != nil {
		return errors.New(err.Error())
	}

	_, err = stmt.Exec(todoId)
	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func DeleteTodo(todoId int64) error {

	query := `DELETE FROM todos WHERE id = ?`

	stmt, err := database.DB.Prepare(query)

	if err != nil {
		return errors.New(err.Error())
	}

	defer stmt.Close()

	_, err = stmt.Exec(todoId)
	if err != nil {
		return errors.New(err.Error())
	}

	return nil

}
