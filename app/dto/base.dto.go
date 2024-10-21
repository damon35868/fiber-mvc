package dto

type (
	PageReqDto struct {
		Page     int         `json:"page" validate:"required,min=1"`
		PageSize int         `json:"pageSize" validate:"required,min=1"`
		Where    interface{} `json:"where"`
	}
)

type (
	PageRespDto struct {
		Items       interface{} `json:"items"`
		TotalCount  int64       `json:"totalCount"`
		HasNextPage bool        `json:"hasNextPage"`
	}

	BaseRespDto struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
)
