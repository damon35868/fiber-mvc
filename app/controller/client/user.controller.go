package clientController

import (
	"fiber-mvc/app/common/utils"
	"fiber-mvc/app/dto"

	"github.com/damon35868/finalx-gofiber/common"
	"github.com/gofiber/fiber/v2"
)

func (s *ClientController) Login(c *fiber.Ctx) error {
	var loginUser dto.LoginReq
	err := c.BodyParser(&loginUser)
	err = utils.Validate.Struct(loginUser)
	if err != nil {
		return common.HttpException(c, fiber.StatusNotFound, err.Error())
	}

	return s.Service.Client.Login(c, &loginUser)
}

func (s *ClientController) Info(c *fiber.Ctx) error {
	return s.Service.Client.Info(c)

}
