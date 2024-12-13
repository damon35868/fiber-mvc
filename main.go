package main

import (
	"fiber-mvc/app/service"
	"fiber-mvc/config"
	"fiber-mvc/schedule"

	_ "fiber-mvc/docs"
	"fiber-mvc/router"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// Swagger
// @title		Title
// @version		1.0
// @description	this is description.
func main() {
	app := fiber.New(config.NewApp())
	app.Use(cors.New())
	app.Use(recover.New())

	serviceInstance := service.New(config.NewDB(), config.NewRedis())
	schedule.Boot(serviceInstance)
	router.Boot(app, serviceInstance)
	app.Listen(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))
}
