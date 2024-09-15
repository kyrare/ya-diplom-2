package postgres

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/kyrare/ya-diplom-2/internal/domain/entities"
)

type UserSecretRepository struct {
	db       *sql.DB
	userRepo *UserRepository
}

func NewPostgresUserSecretRepository(db *sql.DB, userRepo *UserRepository) *UserSecretRepository {
	return &UserSecretRepository{
		db:       db,
		userRepo: userRepo,
	}
}

func (r *UserSecretRepository) Create(secret *entities.ValidatedUserSecret) (*entities.UserSecret, error) {
	row := r.db.QueryRow(
		"insert into user_secrets (id, user_id, type, name, file, file_size, created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7, $8)",
		secret.Id,
		secret.User.Id,
		secret.Type,
		secret.Name,
		secret.File,
		secret.FileSize,
		secret.CreatedAt,
		secret.UpdatedAt,
	)

	if err := row.Err(); err != nil {
		return nil, err
	}

	storedSecret, err := r.FindById(secret.Id)
	if err != nil {
		return nil, err
	}

	if storedSecret == nil {
		return nil, errors.New("не удалось найти секрет после создания")
	}

	return storedSecret, nil
}

func (r *UserSecretRepository) FindById(id uuid.UUID) (*entities.UserSecret, error) {
	var secret entities.UserSecret
	row := r.db.QueryRow("select id, user_id, type, name, file, file_size, created_at, updated_at from user_secrets where id = $1", id)

	var userId uuid.UUID
	err := row.Scan(&secret.Id, &userId, &secret.Type, &secret.Name, &secret.File, &secret.FileSize, &secret.CreatedAt, &secret.UpdatedAt)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	user, err := r.userRepo.FindById(userId)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("не найден пользователя для секрета")
	}

	secret.User = user

	return &secret, nil
}
