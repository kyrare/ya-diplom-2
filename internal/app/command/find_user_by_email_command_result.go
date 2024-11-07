package command

import "github.com/kyrare/ya-diplom-2/internal/domain/entities"

type FindUserByLoginCommandResult struct {
	User *entities.User
}
