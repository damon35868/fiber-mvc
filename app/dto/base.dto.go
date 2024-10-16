package dto

type (
	PageReqDto struct {
		Page     int         `json:"page"`
		PageSize int         `json:"pageSize"`
		Where    interface{} `json:"where"`
	}
)

type (
	PageRespDto struct {
		Items       []interface{} `json:"items"`
		TotalCount  int           `json:"totalCount"`
		HasNextPage bool          `json:"hasNextPage"`
	}

	BaseRespDto struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
)
