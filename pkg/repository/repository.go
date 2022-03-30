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
	Delete(userID, listID int) error
}

type TodoItem interface {
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
	}
}
