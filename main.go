package main

import (
	"fiber-mvc/app/service"
	"fiber-mvc/app/sqlc"
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
	app.Use(recover.New()) //异常处理

	dbConn, err := config.NewDB()
	if err != nil {
		fmt.Println("---数据库连接失败---", err.Error())
		return
	}
	dbInstance := sqlc.New(dbConn)

	// redis
	redisInstance := config.NewRedis()
	if redisInstance == nil {
		fmt.Println("---Redis连接失败---")
		return
	}

	// 挂载持久化实例
	server := service.New(dbInstance, redisInstance)
	// router
	router.ClientRegister(app, server)
	router.AdminRegister(app, server)

	// doc
	app.Get("/swagger/*", swagger.New(config.NewDoc()))

	app.Listen(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))
}
