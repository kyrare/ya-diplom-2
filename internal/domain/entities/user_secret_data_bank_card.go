package entities

import (
	"encoding/json"
	"errors"
	"fmt"
)

type UserSecretDataBankCard struct {
	Number string `json:"number"`
	Month  int64  `json:"month"`
	Year   int64  `json:"year"`
	Cvv    int64  `json:"cvv"`
}

var _ UserSecretData = &UserSecretDataBankCard{}

func NewUserSecretBankCard(number string, month, year, cvv int64) *UserSecretDataBankCard {
	return &UserSecretDataBankCard{
		Number: number,
		Month:  month,
		Year:   year,
		Cvv:    cvv,
	}
}

func newUserSecretBankCardFromData(data []byte) (*UserSecretDataBankCard, error) {
	secretData := new(UserSecretDataBankCard)
	err := json.Unmarshal(data, secretData)
	if err != nil {
		return nil, err
	}
	return secretData, nil
}

func (d *UserSecretDataBankCard) GetType() UserSecretType {
	return UserSecretBankCardType
}

func (d *UserSecretDataBankCard) GetData() ([]byte, error) {
	return json.Marshal(d)
}

func (d *UserSecretDataBankCard) validate() error {
	if len(d.Number) != 16 {
		return errors.New("card number must be 16 characters long")
	}
	for _, digit := range d.Number {
		if digit < '0' || digit > '9' {
			return fmt.Errorf("invalid character in card number: %c", digit)
		}
	}

	if d.Month < 0 || d.Month > 12 {
		return errors.New("card month must be between 1 and 12")
	}

	if d.Year < 2000 || d.Year > 3000 {
		return errors.New("card year must be between 2000 and 3000")
	}

	if d.Cvv < 0 || d.Cvv > 999 {
		return errors.New("card year must be between 000 and 999")
	}

	return nil
}
