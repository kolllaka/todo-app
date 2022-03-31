package repository

import (
	"fmt"
	"strings"

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

func (r *TodoItemMySql) Update(listID, itemID int, updateInput todo.UpdateItemInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)

	if updateInput.Title != nil {
		setValues = append(setValues, "title = ?")
		args = append(args, *updateInput.Title)
	}

	if updateInput.Description != nil {
		setValues = append(setValues, "description = ?")
		args = append(args, *updateInput.Description)
	}

	if updateInput.Done != nil {
		setValues = append(setValues, "done = ?")
		args = append(args, *updateInput.Done)
	}

	setStmt := strings.Join(setValues, ", ")

	// UPDATE todo_items ti, lists_items li SET %s WHERE ti.id = li.item_id AND li.item_id = 1 AND li.list_id = 1
	stmt := fmt.Sprintf("UPDATE %s ti, %s li SET %s WHERE ti.id = li.item_id AND li.item_id = ? AND li.list_id = ?",
		todoItemsTable, listsItemsTable, setStmt)
	args = append(args, itemID, listID)

	_, err := r.db.Exec(stmt, args...)

	return err
}

func (r *TodoItemMySql) Delete(listID, itemID int) error {
	// DELETE ti, li FROM todo_items ti INNER JOIN lists_items li WHERE ti.id = li.item_id AND ti.id = 1 AND li.list_id = 1
	stmt := fmt.Sprintf("DELETE ti, li FROM %s ti INNER JOIN %s li WHERE ti.id = li.item_id AND ti.id = ? AND li.list_id = ?",
		todoItemsTable, listsItemsTable)

	_, err := r.db.Exec(stmt, itemID, listID)

	return err
}
