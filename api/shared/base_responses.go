package shared

import (
	"net/http"
	"time"
)

type BaseResponse struct {
	Success        bool        `json:"success" example:"true"`
	RequestDateUtc time.Time   `json:"request_date_utc" example:"2023-11-09T21:21:46+00:00"`
	Data           interface{} `json:"data"`
}

func SuccessResponse(data interface{}) (int, BaseResponse) {
	return http.StatusOK, BaseResponse{
		Success:        true,
		RequestDateUtc: time.Now().UTC(),
		Data:           data,
	}
}
