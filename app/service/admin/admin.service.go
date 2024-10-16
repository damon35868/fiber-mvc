package adminService

import "fiber-mvc/app/common/types"

type AdminService struct {
	Storage *types.Storage
}

func New(storage *types.Storage) *AdminService {
	return &AdminService{
		Storage: storage,
	}
}
