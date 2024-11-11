package adminService

import (
	"encoding/json"
	"fiber-mvc/app/common/utils"
	"fiber-mvc/app/dto"
	"fiber-mvc/app/sqlc"
	"fmt"
	"os"
	"time"

	"github.com/damon35868/finalx-gofiber/common"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

func (s *AdminService) GetUsers(c *fiber.Ctx, req dto.PageReqDto) error {
	limit, Offset := utils.FormatPage(req)

	totalCount, err := s.Storage.Repository.CountUsers(c.Context())
	users, err := s.Storage.Repository.ListUsers(c.Context(), sqlc.ListUsersParams{
		Limit:  limit,
		Offset: Offset,
	})
	if err != nil {
		return common.HttpException(c, fiber.StatusBadRequest, err.Error())
	}

	userList := []dto.UserResp{}
	copier.Copy(&userList, &users)

	return common.Response(c, &dto.PageRespDto{
		Items:       &userList,
		TotalCount:  totalCount,
		HasNextPage: utils.HasNextPage(req.Page, req.PageSize, int(totalCount)),
	})
}

// 事务示例
func (s *AdminService) CreateUser(c *fiber.Ctx, user *dto.UserReq) error {
	tx, err := s.Storage.DB.Begin()
	if err != nil {
		return fiber.NewError(fiber.StatusBadGateway, "事务开始失败")
	}
	defer tx.Rollback()

	qtx := s.Storage.Repository.WithTx(tx)
	rel, err := qtx.CreateUser(c.Context(), sqlc.CreateUserParams{
		Nickname: "test",
		Age:      18,
		Gender:   1,
		Avatar:   "",
		Password: "123",
	})
	if err != nil {
		return fiber.NewError(fiber.StatusBadGateway, "创建数据失败")
	}

	id, _ := rel.LastInsertId()

	if _, err := qtx.UpdateUser(c.Context(), sqlc.UpdateUserParams{
		ID:       int(id),
		Nickname: "这是更改的内容",
		Age:      60,
		Gender:   0,
		Avatar:   "更改的头像",
		Password: "888888",
	}); err != nil {
		return fiber.NewError(fiber.StatusBadGateway, "更新数据失败")
	}

	tx.Commit()
	return common.Response(c, true)
}

func (s *AdminService) Login(c *fiber.Ctx, req *dto.LoginReq) error {
	user, err := s.Storage.Repository.GetUserByNickName(c.Context(), req.Name)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "该用户不存在")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "密码错误")
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
		"admin":  true,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	})

	token, err := t.SignedString([]byte(os.Getenv("JWT_ADMIN_SECRET")))
	if err != nil {
		return common.HttpException(c, fiber.StatusInternalServerError, err.Error())
	}

	userData := &dto.UserResp{}
	err = copier.Copy(&userData, &user)
	if err != nil {
		return common.HttpException(c, fiber.StatusBadRequest, err.Error())
	}

	return common.Response(c, &fiber.Map{
		"token": token,
		"user":  &userData,
	})
}

func (s *AdminService) GetUserInfo(c *fiber.Ctx) error {
	resp := dto.UserResp{}
	userId := common.GetTokenUserId(c)
	cacheKey := fmt.Sprintf("user:info:%d", userId)

	cache, _ := s.Storage.Cache.Get(cacheKey)
	res := s.Storage.Cache.Conn().SetNX(c.Context(), cacheKey, "1", time.Minute*1)

	if status, _ := res.Result(); !status {
		return fiber.NewError(fiber.StatusBadRequest, "来晚了，已经被抢走了")
	}

	if cache != nil {
		_ = json.Unmarshal(cache, &resp)
		return common.Response(c, &resp)
	}

	user, err := s.Storage.Repository.GetUserById(c.Context(), int(userId))
	data, err := json.Marshal(user)

	err = s.Storage.Cache.Set(cacheKey, data, time.Minute*5)
	s.Storage.Cache.Delete(cacheKey)

	if err != nil {
		return err
	}

	copier.Copy(&resp, &user)
	return common.Response(c, &resp)
}
