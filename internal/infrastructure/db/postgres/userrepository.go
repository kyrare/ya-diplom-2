package postgres

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/kyrare/ya-diplom-2/internal/domain/entities"
)

type UserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(user *entities.ValidatedUser) (*entities.User, error) {
	row := r.db.QueryRow(
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

	createdUser, err := r.FindById(user.Id)
	if err != nil {
		return nil, err
	}

	if createdUser == nil {
		return nil, fmt.Errorf("не удалось найти пользователя после создания")
	}

	return createdUser, nil
}

func (r *UserRepository) FindById(id uuid.UUID) (*entities.User, error) {
	var user entities.User
	row := r.db.QueryRow("select id, login, password, created_at, updated_at from users where id = $1", id)

	err := row.Scan(&user.Id, &user.Login, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return &user, nil
}

func (r *UserRepository) FindByLogin(login string) (*entities.User, error) {
	var user entities.User
	row := r.db.QueryRow("select id, login, password, created_at, updated_at from users where login = $1", login)

	err := row.Scan(&user.Id, &user.Login, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return &user, nil
}

func (r *UserRepository) Delete(id uuid.UUID) error {
	row := r.db.QueryRow("delete from users where id = $1", id)

	return row.Err()
}
