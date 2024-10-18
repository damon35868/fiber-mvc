package adminControllrt

import (
	"fiber-mvc/app/common/utils"
	"fiber-mvc/app/dto"

	"github.com/damon35868/finalx-gofiber/common"
	"github.com/gofiber/fiber/v2"
)

func (s *AdminController) GetBlogs(c *fiber.Ctx) error {
	var pageDto dto.PageReqDto
	err := c.QueryParser(&pageDto)
	if err != nil {
		return common.HttpException(c, fiber.StatusBadRequest, err.Error())
	}
	utils.Validate.Struct(pageDto)

	return s.Service.Admin.GetBlogs(c, pageDto)
}
