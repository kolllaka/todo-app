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
	var id int
	stmt := fmt.Sprintf("INSERT INTO %s (name, username, pasword_hash) VALUES (?, ?, ?)", usersTable)

	row := r.db.QueryRow(stmt, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
