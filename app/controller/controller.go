package controller

import (
	adminController "fiber-mvc/app/controller/admin"
	clientController "fiber-mvc/app/controller/client"
	"fiber-mvc/app/service"
)

type Controller struct {
	AdminController  *adminController.AdminController
	ClientController *clientController.ClientController
}

func New(service *service.Service) *Controller {

	return &Controller{
		AdminController:  adminController.New(service),
		ClientController: clientController.New(service),
	}
}
