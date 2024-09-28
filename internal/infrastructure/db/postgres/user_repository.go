package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/kyrare/ya-diplom-2/internal/domain/entities"
	"github.com/lib/pq"
)

type UserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(ctx context.Context, user *entities.ValidatedUser) (*entities.User, error) {
	row := r.db.QueryRowContext(
		ctx,
		"insert into users (id, login, password, created_at, updated_at) values ($1, $2, $3, $4, $5)",
		user.Id,
		user.Login,
		user.Password,
		user.CreatedAt,
		user.UpdatedAt,
	)

	if err := row.Err(); err != nil {
		return nil, err
	}

	createdUser, err := r.FindById(ctx, user.Id)
	if err != nil {
		return nil, err
	}

	if createdUser == nil {
		return nil, fmt.Errorf("не удалось найти пользователя после создания")
	}

	return createdUser, nil
}

func (r *UserRepository) FindById(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	var user entities.User
	row := r.db.QueryRowContext(ctx, "select id, login, password, created_at, updated_at from users where id = $1", id)

	err := row.Scan(&user.Id, &user.Login, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return &user, nil
}

func (r *UserRepository) FindByIDs(ctx context.Context, IDs []uuid.UUID) ([]*entities.User, error) {
	usersIDs := make([]string, len(IDs))
	for i, id := range IDs {
		usersIDs[i] = id.String()
	}

	rows, err := r.db.QueryContext(ctx, "select id, login, password, created_at, updated_at from users where id = any($1)", pq.Array(usersIDs))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]*entities.User, 0)

	for rows.Next() {
		var user entities.User
		err = rows.Scan(&user.Id, &user.Login, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}

		result = append(result, &user)
	}

	return result, nil
}

func (r *UserRepository) FindByLogin(ctx context.Context, login string) (*entities.User, error) {
	var user entities.User
	row := r.db.QueryRowContext(ctx, "select id, login, password, created_at, updated_at from users where login = $1", login)

	err := row.Scan(&user.Id, &user.Login, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return &user, nil
}

func (r *UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	row := r.db.QueryRowContext(ctx, "delete from users where id = $1", id)

	return row.Err()
}
