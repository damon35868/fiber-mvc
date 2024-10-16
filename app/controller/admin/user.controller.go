package adminControllrt

import (
	"fiber-mvc/app/common/utils"
	"fiber-mvc/app/dto"

	"github.com/damon35868/finalx-gofiber/common"
	"github.com/gofiber/fiber/v2"
)

func (s *AdminController) GetUsers(c *fiber.Ctx) error {
	return s.Service.Admin.GetUsers(c)
}

func (s *AdminController) CreateUser(c *fiber.Ctx) error {
	var user dto.UserReq
	err := c.BodyParser(&user)
	if err != nil {
		return common.HttpException(c, fiber.StatusBadRequest, err.Error())
	}
	utils.Validate.Struct(user)

	return s.Service.Admin.CreateUser(c, &user)
}

// @Router		/user/login [POST]
// @Tags		admin - 用户相关
// @Summary		admin - 登陆接口
// @Param		id	path int true "路径ID"
// @Param 		email body dto.UserResp true "请求体"
// @Success		200	{object} dto.UserResp
func (s *AdminController) Login(c *fiber.Ctx) error {
	var req *dto.LoginReq
	err := c.BodyParser(&req)
	err = utils.Validate.Struct(req)
	if err != nil {
		return common.HttpException(c, fiber.StatusBadRequest, err.Error())
	}
	return s.Service.Admin.Login(c, req)

}

// @Router		/user/{id} [GET]
// @Tags		admin - 用户相关
// @Summary		admin - 用户详情
// @Param		id	path int true "路径ID"
// @Success		200	{object} dto.UserResp
func (s *AdminController) GetUserInfo(c *fiber.Ctx) error {
	return s.Service.Admin.GetUserInfo(c)

}
