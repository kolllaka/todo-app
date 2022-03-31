package repository

import (
	"fmt"

	"github.com/KoLLlaka/todo-app/internal/todo"
	"github.com/jmoiron/sqlx"
)

type TodoItemMySql struct {
	db *sqlx.DB
}

func NewTodoItemMySql(db *sqlx.DB) *TodoItemMySql {
	return &TodoItemMySql{db: db}
}

func (r *TodoItemMySql) Create(listID int, input todo.TodoItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	// INSERT INTO todo_items (title, description, done) VALUES ("test", "test", true);
	createItemStmt := fmt.Sprintf("INSERT INTO %s (title, description, done) VALUES (?, ?, ?)", todoItemsTable)
	result, err := r.db.Exec(createItemStmt, input.Title, input.Description, input.Done)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	itemID, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	// INSERT INTO lists_items (user_id, item_id) VALUES (1, 1);
	createdUsersListStmt := fmt.Sprintf("INSERT INTO %s (list_id, item_id) VALUES (?, ?)", listsItemsTable)
	_, err = r.db.Exec(createdUsersListStmt, listID, itemID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return int(itemID), nil
}
