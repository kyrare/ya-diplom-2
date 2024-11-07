package entities

import (
	"errors"
	"time"
	"unicode/utf8"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID
	Login     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(login string, password string) *User {
	return &User{
		Id:        uuid.New(),
		Login:     login,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (u *User) validate() error {
	if u.Login == "" {
		return errors.New("login must not be empty")
	}

	if utf8.RuneCountInString(u.Login) < 3 {
		return errors.New("login must be 6 characters long")
	}

	if utf8.RuneCountInString(u.Login) > 255 {
		return errors.New("login must be shorter than 255 characters")
	}

	if u.Password == "" {
		return errors.New("password must not be empty")
	}

	if utf8.RuneCountInString(u.Password) < 6 {
		return errors.New("password must be 6 characters long")
	}

	if utf8.RuneCountInString(u.Password) > 255 {
		return errors.New("password must be shorter than 255 characters")
	}

	if u.CreatedAt.After(u.UpdatedAt) {
		return errors.New("created_at must be before updated_at")
	}

	return nil
}
