package utils

import (
	"fiber-mvc/app/dto"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func FormatPage(req dto.PageReqDto) (int32, int32) {
	return int32(req.PageSize), int32(req.Page-1) * int32(req.PageSize)
}
