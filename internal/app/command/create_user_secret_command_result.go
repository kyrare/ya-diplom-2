package command

import (
	"github.com/kyrare/ya-diplom-2/internal/domain/entities"
)

type CreateUserSecretCommandResult struct {
	Secret *entities.UserSecret
}
