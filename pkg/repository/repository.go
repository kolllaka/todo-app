package repository

import (
	"github.com/KoLLlaka/todo-app/internal/todo"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUserID(username, password string) (todo.User, error)
}

type TodoList interface {
	Create(userID int, list todo.TodoList) (int, error)
	GetAll(userID int) ([]todo.TodoList, error)
	GetByID(userID, listID int) (todo.TodoList, error)
	Update(userID, listID int, updateInput todo.UpdateListInput) error
	Delete(userID, listID int) error
}

type TodoItem interface {
	Create(listID int, input todo.TodoItem) (int, error)
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthMySql(db),
		TodoList:      NewTodoListMySql(db),
		TodoItem:      NewTodoItemMySql(db),
	}
}
