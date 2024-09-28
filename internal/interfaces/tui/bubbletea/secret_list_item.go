package bubbletea

import (
	"fmt"
	"strconv"

	"github.com/kyrare/ya-diplom-2/internal/domain/entities"
)

type SecretListItem struct {
	entities.UserSecret
}

func (i SecretListItem) Title() string       { return i.Name }
func (i SecretListItem) FilterValue() string { return i.Name }
func (i SecretListItem) Description() string {
	switch i.Type {
	case entities.UserSecretPasswordType:
		data := (*i.Data).(*entities.UserSecretDataPassword)

		return fmt.Sprintf("Login: %s Password: %s", data.Login, data.Password)

	case entities.UserSecretBankCardType:
		data := (*i.Data).(*entities.UserSecretDataBankCard)

		exp := strconv.FormatInt(data.Month, 10) + "/" + strconv.FormatInt(data.Year-2000, 10)

		return fmt.Sprintf("exp: %s cvv: %d", exp, data.Cvv)

	case entities.UserSecretTextType:
		data := (*i.Data).(*entities.UserSecretDataText)
		runes := []rune(data.Text)
		if len(runes) > 40 {
			runes = runes[:40]
			runes = append(runes, '.', '.', '.')
		}

		return string(runes)
	}

	return ""
}
