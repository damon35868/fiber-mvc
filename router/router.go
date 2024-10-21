package router

import (
	"fiber-mvc/app/service"

	"github.com/gofiber/fiber/v2"
)

func Boot(app *fiber.App, server *service.Service) {
	ClientRegister(app, server)
	AdminRegister(app, server)
}
