package adminControllrt

import (
	"fiber-mvc/app/common/utils"
	"fiber-mvc/app/dto"

	"github.com/damon35868/finalx-gofiber/common"
	"github.com/gofiber/fiber/v2"
)

func (s *AdminController) GetBlogs(c *fiber.Ctx) error {
	var pageDto dto.BlogListDto
	err := c.BodyParser(&pageDto)
	if err != nil {
		return common.HttpException(c, fiber.StatusBadRequest, err.Error())
	}
	if err = utils.Validate.Struct(&pageDto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return s.Service.Admin.GetBlogs(c, pageDto)
}
