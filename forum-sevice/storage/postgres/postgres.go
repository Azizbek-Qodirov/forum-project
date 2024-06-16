package postgres

import (
	"database/sql"
	"fmt"

	"reservation-service/config"
	"reservation-service/config/logger"

	_ "github.com/lib/pq"
)

type Storage struct {
	Db     *sql.DB
	Logger *logger.Logger
}

func NewPostgresStorage(config config.Config, logger *logger.Logger) (*Storage, error) {
	conn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%d sslmode=disable",
		config.DB_HOST, config.DB_USER, config.DB_NAME, config.DB_PASSWORD, config.DB_PORT)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Storage{
		Db: db,
	}, nil
}
