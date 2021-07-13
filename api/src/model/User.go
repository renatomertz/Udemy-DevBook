package model

import (
	"errors"
	"strings"
	"time"
)

//User
type User struct {
	ID       uint64    `json:"id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Password string    `json:"password,omitempty"`
	Created  time.Time `json:"created,omitempty"`
}

// Prepare
func (user *User) Prepare() error {
	user.formatter()
	if err := user.validate(); err != nil {
		return err
	}
	return nil
}

func (user *User) validate() error {
	if user.Name == "" {
		return errors.New("The name is mandatory")
	}

	if user.Nick == "" {
		return errors.New("The nick is mandatory")
	}

	if user.Email == "" {
		return errors.New("The email is mandatory")
	}

	if user.Password == "" {
		return errors.New("The password is mandatory")
	}
	return nil
}

func (user *User) formatter() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}
