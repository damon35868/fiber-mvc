package types

import (
	"fiber-mvc/app/sqlc"

	"github.com/gofiber/storage/redis/v3"
)

type Storage struct {
	Repository *sqlc.Queries
	Cache      *redis.Storage
}
