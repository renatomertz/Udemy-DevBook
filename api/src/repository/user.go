package repository

import (
	"api/src/model"
	"database/sql"
)

// Users (Repository)
type Users struct {
	db *sql.DB
}

// Create repository
func NewUsersRepository(db *sql.DB) *Users {
	return &Users{db}
}

//Create user
func (repository Users) Create(user model.User) (uint64, error) {
	statement, err := repository.db.Prepare("INSERT INTO user (name, nick, email, password) VALUES (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	resultado, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastInsertedId, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertedId), nil
}
