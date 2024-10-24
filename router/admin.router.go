package router

import (
	"fiber-mvc/app/common/constant"
	"fiber-mvc/app/controller"
	"fiber-mvc/app/service"
	"os"

	"github.com/damon35868/finalx-gofiber/common"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func AdminRegister(app *fiber.App, server *service.Service) {
	handler := controller.New(server)
	admin := app.Group(constant.Admin)
	// 不走JWT鉴权
	admin.Post("/user/login", handler.AdminController.Login)

	// JWT
	admin.Use(jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(os.Getenv("JWT_ADMIN_SECRET"))},
		ErrorHandler: common.JWTErrorHandler,
	}))

	// user
	adminUserApi := admin.Group("/user")
	{
		adminUserApi.Get("/list", handler.AdminController.GetUsers)
		adminUserApi.Post("/new", handler.AdminController.CreateUser)
		adminUserApi.Get("/", handler.AdminController.GetUserInfo)
	}

	// blog
	blogApi := admin.Group("/blog")
	{
		blogApi.Post("/list", handler.AdminController.GetBlogs)
	}
}
