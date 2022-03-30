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

func (r *TodoListMySql) GetAll(userID int) ([]todo.TodoList, error) {
	var lists []todo.TodoList
	// SELECT tl.id, title, description FROM todo_lists tl INNER JOIN users_lists ul ON tl.id = ul.list_id WHERE ul.user_id = 1
	stmt := fmt.Sprintf("SELECT tl.id, title, description FROM %s tl INNER JOIN %s ul ON tl.id = ul.list_id WHERE ul.user_id = ?",
		todoListsTable, usersListsTable)

	rows, err := r.db.Query(stmt, userID)
	if err != nil {
		return []todo.TodoList{}, err
	}

	for rows.Next() {
		var list todo.TodoList
		if err := rows.Scan(&list.Id, &list.Title, &list.Description); err != nil {
			return []todo.TodoList{}, err
		}

		lists = append(lists, list)
	}

	return lists, nil
}

func (r *TodoListMySql) GetByID(userID, listID int) (todo.TodoList, error) {
	var list todo.TodoList
	// SELECT tl.id, title, description FROM todo_lists tl INNER JOIN users_lists ul ON tl.id = ul.list_id WHERE ul.user_id = 1 AND ul.list_id = 1
	stmt := fmt.Sprintf("SELECT tl.id, title, description FROM %s tl INNER JOIN %s ul ON tl.id = ul.list_id WHERE ul.user_id = ? AND ul.list_id = ?",
		todoListsTable, usersListsTable)

	row := r.db.QueryRow(stmt, userID, listID)
	if err := row.Scan(&list.Id, &list.Title, &list.Description); err != nil {
		return todo.TodoList{}, err
	}

	return list, nil
}

func (r *TodoListMySql) Delete(userID, listID int) error {
	// DELETE tl, ul FROM todo_lists tl INNER JOIN users_lists ul WHERE tl.id = ul.list_id AND tl.id = 1 AND ul.user_id = 1
	stmt := fmt.Sprintf("DELETE tl, ul FROM %s tl INNER JOIN %s ul WHERE tl.id = ul.list_id AND ul.user_id = ? AND ul.list_id = ?",
		todoListsTable, usersListsTable)

	_, err := r.db.Exec(stmt, userID, listID)

	return err
}
