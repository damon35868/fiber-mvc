package clientService

import (
	"encoding/json"
	"fiber-mvc/app/dto"
	"fmt"
	"os"
	"time"

	"github.com/damon35868/finalx-gofiber/common"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

func (s *ClientService) Login(c *fiber.Ctx, req *dto.LoginReq) error {
	user, _ := s.Storage.Repository.GetUserByNickName(c.Context(), req.Name)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return common.HttpException(c, fiber.StatusBadRequest, "密码错误")
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
		"admin":  true,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	})

	token, err := t.SignedString([]byte(os.Getenv("JWT_CLIENT_SECRET")))
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

func (s *ClientService) Info(c *fiber.Ctx) error {
	user := &dto.UserResp{}
	userId := common.GetTokenUserId(c)

	cacheKey := fmt.Sprintf("user:info:%d", userId)

	d, _ := s.Storage.Cache.Get(cacheKey)
	if d != nil {
		_ = json.Unmarshal(d, &user)
		return common.Response(c, &user)
	}

	db, err := s.Storage.Repository.GetUserById(c.Context(), int(userId))
	data, err := json.Marshal(db)
	err = s.Storage.Cache.Set(cacheKey, data, time.Minute*5)
	if err != nil {
		return common.HttpException(c, fiber.StatusBadRequest, err.Error())
	}

	err = copier.Copy(&user, &db)
	if err != nil {
		return common.HttpException(c, fiber.StatusFailedDependency, err.Error())
	}
	return common.Response(c, &user)
}
