package service

import (
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

func New(db *sqlc.Queries, cache *redis.Storage) *Service {
	storage := &types.Storage{
		Repository: db,
		Cache:      cache,
	}
	return &Service{
		Admin:  adminService.New(storage),
		Client: clientService.New(storage),
	}
}
