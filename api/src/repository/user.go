package repository

import (
	"api/src/model"
	"database/sql"
	"fmt"
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

// Find
func (repository Users) Find(nameOrNick string) ([]model.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)
	lines, err := repository.db.Query(
		"SELECT id, name, nick, email, created FROM user WHERE name LIKE ? OR nick LIKE ? ",
		nameOrNick,
		nameOrNick,
	)
	if err != nil {
		return nil, err
	}

	defer lines.Close()
	var users []model.User

	for lines.Next() {
		var user model.User
		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.Created,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// Find by user ID
func (repository Users) FindById(ID uint64) (model.User, error) {
	lines, err := repository.db.Query(
		"SELECT id, name, nick, email, created FROM user WHERE id = ?",
		ID,
	)

	if err != nil {
		return model.User{}, err
	}
	defer lines.Close()

	var user model.User

	if lines.Next() {
		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.Created,
		); err != nil {
			return model.User{}, err
		}
	}
	return user, nil
}
