package command

import "github.com/kyrare/ya-diplom-2/internal/domain/entities"

type FindUserByIdCommandResult struct {
	User *entities.User
}
