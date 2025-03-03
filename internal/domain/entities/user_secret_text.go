package entities

import "errors"

type UserSecretDataText struct {
	Text string
}

var _ UserSecretData = &UserSecretDataText{}

func NewUserSecretText(text string) *UserSecretDataText {
	return &UserSecretDataText{
		Text: text,
	}
}

func newUserSecretTextFromData(data []byte) (*UserSecretDataText, error) {
	secretData := NewUserSecretText(string(data))

	return secretData, nil
}

func (d *UserSecretDataText) GetType() UserSecretType {
	return UserSecretTextType
}

func (d *UserSecretDataText) GetData() ([]byte, error) {
	return []byte(d.Text), nil
}

func (d *UserSecretDataText) validate() error {
	if d.Text == "" {
		return errors.New("text must not be empty")
	}

	return nil
}
