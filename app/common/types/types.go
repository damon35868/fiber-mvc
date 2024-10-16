package types

import (
	"database/sql"
	"fiber-mvc/app/sqlc"

	"github.com/gofiber/storage/redis/v3"
)

type Storage struct {
	DB         *sql.DB
	Repository *sqlc.Queries
	Cache      *redis.Storage
}
