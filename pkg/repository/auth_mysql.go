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
	// INSERT INTO users (name, username, password_hash) VALUES ("test", "test", "test")
	stmt := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES (?, ?, ?);", usersTable)

	result, err := r.db.Exec(stmt, user.Name, user.Username, user.Password)
	if err != nil {
		return 0, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(lastID), nil
}

func (r *AuthMySql) GetUserID(username, password string) (todo.User, error) {
	var user todo.User
	// SELECT * FROM users WHERE username = "test" AND password_hash = "test"
	stmt := fmt.Sprintf("SELECT id FROM %s WHERE username = ? AND password_hash = ?", usersTable)

	err := r.db.Get(&user, stmt, username, password)

	return user, err
}
