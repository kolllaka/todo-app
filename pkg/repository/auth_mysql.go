package repository

import (
	"fmt"

	"github.com/KoLLlaka/todo-app/internal/todo"

	"github.com/jmoiron/sqlx"
)

type AuthMySql struct {
	db *sqlx.DB
}

func NewAuthMySql(db *sqlx.DB) *AuthMySql {
	return &AuthMySql{db: db}
}

func (r *AuthMySql) CreateUser(user todo.User) (int, error) {
	stmt := fmt.Sprintf("INSERT INTO %s (name, username, pasword_hash) VALUES (?, ?, ?);", usersTable)

	result, err := r.db.Exec(stmt, user.Name, user.Username, user.Password)
	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(lastId), nil
}
