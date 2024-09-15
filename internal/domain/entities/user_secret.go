package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type UserSecretType string

const (
	UserSecretPasswordType = UserSecretType("password")
	UserSecretBankCardType = UserSecretType("bank_card")
	UserSecretTextType     = UserSecretType("text")
	UserSecretFileType     = UserSecretType("file")
)

type UserSecret struct {
	Id        uuid.UUID
	User      *User
	Type      UserSecretType
	Name      string // Название файла которое отображается пользователю
	Data      *UserSecretData
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUserSecret(u *User, t UserSecretType, name string, d *UserSecretData) *UserSecret {
	return &UserSecret{
		Id:        uuid.New(),
		User:      u,
		Type:      t,
		Name:      name,
		Data:      d,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (us *UserSecret) validate() error {
	if us.User == nil {
		return errors.New("секрет должен быть привязан к пользователю")
	}

	if us.Name == "" {
		return errors.New("секрет должен иметь название")
	}

	if us.Data == nil {
		return errors.New("секрет должен содержать данные")
	}

	// todo узнать почему не работает по ссылке
	if err := (*us.Data).validate(); err != nil {
		return err
	}

	return nil
}
