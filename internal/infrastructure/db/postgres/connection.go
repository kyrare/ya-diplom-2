package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/kyrare/ya-diplom-2/internal/app/services"
)

func NewPostgresql(dbName, host, port, user, pass string, logger *services.Logger) (*sql.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, pass, host, port, dbName)

	logger.Info("Create DB connection, dsn - " + dsn)

	return sql.Open("pgx", dsn)
}
