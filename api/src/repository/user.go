package repository

import (
	"api/src/model"
	"database/sql"
)

// Users
type Users struct {
	db *sql.DB
}

// Create repository
func NewUsersRepository(db *sql.DB) *Users {
	return &Users{db}
}

//Create user
func (u Users) Create(user model.User) (uint64, error) {
	return 0, nil
}
