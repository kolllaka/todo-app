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

	// INSERT INTO lists_items (list_id, item_id) VALUES (1, 1);
	createdUsersListStmt := fmt.Sprintf("INSERT INTO %s (list_id, item_id) VALUES (?, ?)", listsItemsTable)
	_, err = r.db.Exec(createdUsersListStmt, listID, itemID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return int(itemID), nil
}

func (r *TodoItemMySql) GetAll(listID int) ([]todo.TodoItem, error) {
	var items []todo.TodoItem
	// SELECT ti.id, title, description, done FROM todo_items ti INNER JOIN lists_items li ON ti.id = li.item_id WHERE li.list_id = 1
	stmt := fmt.Sprintf("SELECT ti.id, title, description, done FROM %s ti INNER JOIN %s li ON ti.id = li.item_id WHERE li.list_id = ?", todoItemsTable, listsItemsTable)

	rows, err := r.db.Query(stmt, listID)
	if err != nil {
		return []todo.TodoItem{}, err
	}

	for rows.Next() {
		var item todo.TodoItem

		if err := rows.Scan(&item.Id, &item.Title, &item.Description, &item.Done); err != nil {
			return []todo.TodoItem{}, err
		}

		items = append(items, item)
	}

	return items, nil
}

func (r *TodoItemMySql) GetByID(listID, itemID int) (todo.TodoItem, error) {
	var item todo.TodoItem
	// SELECT ti.id, title, description, done FROM todo_items ti INNER JOIN lists_items li ON ti.id = li.item_id WHERE li.list_id = 1 AND li.item_id = 1
	stmt := fmt.Sprintf("SELECT ti.id, title, description, done FROM %s ti INNER JOIN %s li ON ti.id = li.item_id WHERE li.list_id = ? AND li.item_id = ?",
		todoItemsTable, listsItemsTable)

	row := r.db.QueryRow(stmt, listID, itemID)
	if err := row.Scan(&item.Id, &item.Title, &item.Description, &item.Done); err != nil {
		return todo.TodoItem{}, err
	}

	return item, nil
}
