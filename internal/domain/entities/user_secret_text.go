package entities

import "errors"

type UserSecretDataText struct {
	text string
}

var _ UserSecretData = &UserSecretDataText{}

func NewUserSecretText(text string) *UserSecretDataText {
	return &UserSecretDataText{
		text: text,
	}
}

func (d *UserSecretDataText) GetType() UserSecretType {
	return UserSecretTextType
}

func (d *UserSecretDataText) GetData() ([]byte, error) {
	return []byte(d.text), nil
}

func (d *UserSecretDataText) validate() error {
	if d.text == "" {
		return errors.New("text must not be empty")
	}

	return nil
}
