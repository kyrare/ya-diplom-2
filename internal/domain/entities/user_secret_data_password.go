package entities

import (
	"encoding/json"
	"errors"
)

type UserSecretDataPassword struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

var _ UserSecretData = &UserSecretDataPassword{}

func NewUserSecretPassword(login, password string) *UserSecretDataPassword {
	return &UserSecretDataPassword{
		Login:    login,
		Password: password,
	}
}

func (d *UserSecretDataPassword) GetType() UserSecretType {
	return UserSecretPasswordType
}

func (d *UserSecretDataPassword) GetData() ([]byte, error) {
	return json.Marshal(d)
}

func (d *UserSecretDataPassword) validate() error {
	if d.Login == "" {
		return errors.New("login must not be empty")
	}

	if d.Password == "" {
		return errors.New("password must not be empty")
	}

	return nil
}
