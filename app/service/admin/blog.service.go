package adminService

import (
	"fiber-mvc/app/common/utils"
	"fiber-mvc/app/dto"
	"fiber-mvc/app/sqlc"

	"github.com/damon35868/finalx-gofiber/common"
	"github.com/gofiber/fiber/v2"
)

func (s *AdminService) GetBlogs(c *fiber.Ctx, req dto.BlogListDto) error {
	limit, offset := utils.FormatPage(*req.PageReqDto)

	blogs, err := s.Storage.Repository.ListBlogs(c.Context(), sqlc.ListBlogsParams{
		Desc:   "%" + req.Where.Desc + "%",
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return common.HttpException(c, fiber.StatusBadRequest, err.Error())
	}
	totalCount, err := s.Storage.Repository.CountBlogs(c.Context(), sqlc.CountBlogsParams{
		Desc: "%" + req.Where.Desc + "%",
	})

	return common.Response(c, &dto.PageRespDto{
		Items:       &blogs,
		TotalCount:  totalCount,
		HasNextPage: utils.HasNextPage(req.Page, req.PageSize, int(totalCount)),
	})
}
