package router

import (
	"fiber-mvc/app/service"
	"fiber-mvc/config"

	"github.com/damon35868/finalx-gofiber/common"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

/**
 * @description: 公共路由
 * @param {*fiber.App} app
 * @return {*}
 */
func commonRegister(app *fiber.App) {
	app.Get("/swagger/*", swagger.New(config.NewDoc()))
	app.Get("/api/liveProbe", func(c *fiber.Ctx) error {
		return common.Response(c, true)
	})
}

func Boot(app *fiber.App, server *service.Service) {
	commonRegister(app)
	ClientRegister(app, server)
	AdminRegister(app, server)
}
