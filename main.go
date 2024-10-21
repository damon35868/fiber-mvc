package main

import (
	"fiber-mvc/app/service"
	"fiber-mvc/config"

	_ "fiber-mvc/docs"
	"fiber-mvc/router"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

func main() {
	app := fiber.New(config.NewApp())
	app.Use(recover.New()) //处理panic

	router.Boot(app, service.New(config.NewDB(), config.NewRedis()))
	// doc
	app.Get("/swagger/*", swagger.New(config.NewDoc()))
	app.Listen(fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT")))
}
