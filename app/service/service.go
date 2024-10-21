package service

import (
	"database/sql"
	"fiber-mvc/app/common/types"
	adminService "fiber-mvc/app/service/admin"
	clientService "fiber-mvc/app/service/client"
	"fiber-mvc/app/sqlc"

	"github.com/gofiber/storage/redis/v3"
)

type Service struct {
	Admin  *adminService.AdminService
	Client *clientService.ClientService
}

func New(db *sql.DB, cache *redis.Storage) *Service {

	storage := &types.Storage{
		DB:         db,
		Repository: sqlc.New(db),
		Cache:      cache,
	}
	return &Service{
		Admin:  adminService.New(storage),
		Client: clientService.New(storage),
	}
}
