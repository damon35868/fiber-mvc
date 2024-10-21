package dto

type (
	BlogListDto struct {
		*PageReqDto
		Where Where `json:"where" validate:"omitempty"`
	}

	Where struct {
		Desc string `json:"desc" validate:"required"`
	}
)
