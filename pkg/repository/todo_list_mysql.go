package repository

import (
	"fmt"

	"github.com/KoLLlaka/todo-app/internal/todo"
	"github.com/jmoiron/sqlx"
)

type TodoListMySql struct {
	db *sqlx.DB
}

func NewTodoListMySql(db *sqlx.DB) *TodoListMySql {
	return &TodoListMySql{db: db}
}

func (r *TodoListMySql) Create(userID int, list todo.TodoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	// INSERT INTO todo_lists (title, description) VALUES ("test", "test");
	createListStmt := fmt.Sprintf("INSERT INTO %s (title, description) VALUES (?, ?)", todoListsTable)
	result, err := tx.Exec(createListStmt, list.Title, list.Description)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	// INSERT INTO users_lists (user_id, list_id) VALUES (1, 1);
	createdUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES (?, ?)", usersListsTable)
	_, err = tx.Exec(createdUsersListQuery, userID, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return int(id), tx.Commit()
}
