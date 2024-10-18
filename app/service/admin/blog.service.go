package adminService

import (
	"fiber-mvc/app/common/utils"
	"fiber-mvc/app/dto"
	"fiber-mvc/app/sqlc"

	"github.com/damon35868/finalx-gofiber/common"
	"github.com/gofiber/fiber/v2"
)

func (s *AdminService) GetBlogs(c *fiber.Ctx, req dto.PageReqDto) error {
	limit, Offset := utils.FormatPage(req)

	blogs, err := s.Storage.Repository.ListBlogs(c.Context(), sqlc.ListBlogsParams{
		Limit:  limit,
		Offset: Offset,
	})
	if err != nil {
		return common.HttpException(c, fiber.StatusBadRequest, err.Error())
	}

	return common.Response(c, &dto.PageRespDto{
		Items: &blogs,
		// TotalCount:  totalCount,
		// HasNextPage: utils.HasNextPage(req.Page, req.PageSize, int(totalCount)),
	})
}
