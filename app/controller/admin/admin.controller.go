package adminControllrt

import "fiber-mvc/app/service"

type AdminController struct {
	Service *service.Service
}

func New(service *service.Service) *AdminController {
	return &AdminController{
		Service: service,
	}
}
