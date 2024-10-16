package clientService

import (
	"fiber-mvc/app/common/types"
)

type ClientService struct {
	Storage *types.Storage
}

func New(storage *types.Storage) *ClientService {
	return &ClientService{
		Storage: storage,
	}
}
