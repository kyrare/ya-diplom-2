package postgres

import (
	"context"
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

func (r *UserSecretRepository) Create(ctx context.Context, secret *entities.ValidatedUserSecret) (*entities.UserSecret, error) {
	row := r.db.QueryRowContext(
		ctx,
		"insert into user_secrets (id, user_id, type, name, created_at, updated_at) values ($1, $2, $3, $4, $5, $6)",
		secret.Id,
		secret.User.Id,
		secret.Type,
		secret.Name,
		secret.CreatedAt,
		secret.UpdatedAt,
	)

	if err := row.Err(); err != nil {
		return nil, err
	}

	storedSecret, err := r.FindById(ctx, secret.Id)
	if err != nil {
		return nil, err
	}

	if storedSecret == nil {
		return nil, errors.New("не удалось найти секрет после создания")
	}

	return storedSecret, nil
}

func (r *UserSecretRepository) FindById(ctx context.Context, id uuid.UUID) (*entities.UserSecret, error) {
	var secret entities.UserSecret
	row := r.db.QueryRowContext(ctx, "select id, user_id, type, name, created_at, updated_at from user_secrets where id = $1", id)

	err := row.Scan(&secret.Id, &secret.UserID, &secret.Type, &secret.Name, &secret.CreatedAt, &secret.UpdatedAt)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	user, err := r.userRepo.FindById(ctx, secret.UserID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("не найден пользователя для секрета")
	}

	secret.User = user

	return &secret, nil
}

func (r *UserSecretRepository) GetAllForUser(ctx context.Context, userID uuid.UUID) ([]*entities.UserSecret, error) {
	rows, err := r.db.QueryContext(ctx, "select id, user_id, type, name, created_at, updated_at from user_secrets WHERE user_id = $1", userID)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	result := make([]*entities.UserSecret, 0)
	usersIDs := make([]uuid.UUID, 0)

	for rows.Next() {
		var secret entities.UserSecret
		err = rows.Scan(&secret.Id, &secret.UserID, &secret.Type, &secret.Name, &secret.CreatedAt, &secret.UpdatedAt)
		if err != nil {
			return nil, err
		}

		result = append(result, &secret)
		usersIDs = append(usersIDs, secret.UserID)
	}

	users, err := r.userRepo.FindByIDs(ctx, usersIDs)
	if err != nil {
		return nil, err
	}

	usersMap := make(map[uuid.UUID]*entities.User, len(users))
	for _, user := range users {
		usersMap[user.Id] = user
	}

	for i, secret := range result {
		if user, ok := usersMap[secret.UserID]; ok {
			result[i].User = user
		}
	}

	return result, nil
}
