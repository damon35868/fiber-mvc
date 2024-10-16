package dto

import "time"

// req
type (
	UserReq struct {
		NickName string `json:"nickname" validate:"required,min=5,max=20"`
		Password string `json:"password" validate:"required,min=4,max=20"`
		Age      int    `json:"age" validate:"required"`
		Gender   int    `json:"gender" validate:"required"`
	}

	LoginReq struct {
		Name     string `json:"name" validate:"required,min=1,max=20"`
		Password string `json:"password" validate:"required,min=4,max=20"`
	}
)

// resp
type (
	UserResp struct {
		ID int `json:"id"`
		// 用户名
		Nickname string `json:"nickname"`
		// 密码
		Password string `json:"-"`
		// 用户性别
		Gender int `json:"gender"`
		// 年龄
		Age int `json:"age"`
		// 头像
		Avatar string `json:"avatar"`
		// 创建时间
		CreatedAt time.Time `json:"createdAt"`
		// 更新时间
		UpdatedAt time.Time `json:"updatedAt"`
	}
)
