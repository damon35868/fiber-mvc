package adminService

import (
	"fiber-mvc/app/common/utils"
	"fiber-mvc/app/dto"
	"fiber-mvc/app/sqlc"
	"fmt"
	"time"

	"github.com/damon35868/finalx-gofiber/common"
	"github.com/gofiber/fiber/v2"
)

type ListBlogsRow1 struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Desc      string    `json:"desc"`
	UserID    int       `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	User      sqlc.User `json:"user"`
}

func (s *AdminService) GetBlogs(c *fiber.Ctx, req dto.BlogListDto) error {
	limit, offset := utils.FormatPage(*req.PageReqDto)

	blogs, err := s.Storage.Repository.ListBlogs(c.Context(), sqlc.ListBlogsParams{
		// Desc:   "%" + req.Where.Desc + "%",
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return common.HttpException(c, fiber.StatusBadRequest, err.Error())
	}
	totalCount, err := s.Storage.Repository.CountBlogs(c.Context(), "%"+req.Where.Desc+"%")

	fmt.Println(blogs)

	var aa = &[]ListBlogsRow1{}

	return common.Response(c, &dto.PageRespDto{
		Items:       &aa,
		TotalCount:  totalCount,
		HasNextPage: utils.HasNextPage(req.Page, req.PageSize, int(totalCount)),
	})
}
