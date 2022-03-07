package db

import (
	"database/sql"

	"github.com/gefion-tech/bfb-user-microservice/internal/config"
	_ "github.com/lib/pq" // database driver
)

// Функция инициализации Postgres БД
func InitPostgres(config *config.DatabaseConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", config.Url)
	if err != nil {
		return nil, err
	}

	// Тест соединения
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
