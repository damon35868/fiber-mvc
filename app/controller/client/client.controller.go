package clientController

import "fiber-mvc/app/service"

type ClientController struct {
	Service *service.Service
}

func New(service *service.Service) *ClientController {
	return &ClientController{
		Service: service,
	}
}
