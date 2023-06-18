package shared

import (
	"net/http"
	"time"
)

type Errors struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details"`
}

type ErrorResponse struct {
	Success        bool      `json:"success" example:"false"`
	RequestDateUtc time.Time `json:"request_date_utc" example:"2023-11-09T21:21:46+00:00"`
	Path           string    `json:"path" example:"/some/api/path"`
	Errors         []Errors  `json:"errors"`
}

func NotFoundError(err error, path string) (int, ErrorResponse) {
	// improvement: assign code and details depending on error received
	return http.StatusNotFound, ErrorResponse{
		Success:        false,
		RequestDateUtc: time.Now().UTC(),
		Path:           path,
		Errors: []Errors{
			{
				Code:    "ERR001",
				Message: "Not found",
				Details: nil,
			},
		},
	}
}

func BadRequestError(err error, path string) (int, ErrorResponse) {
	// improvement: assign code and details depending on error received
	return http.StatusBadRequest, ErrorResponse{
		Success:        false,
		RequestDateUtc: time.Now().UTC(),
		Path:           path,
		Errors: []Errors{
			{
				Code:    "ERR002",
				Message: "Bad request",
				Details: nil,
			},
		},
	}
}
