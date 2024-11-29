package models

type SuccessResponseMessage struct {
	Message string `json:"message"`
}

type ErrorResponseMessage struct {
	Message string `json:"message"`
}

type SuccessResponse struct {
	Data interface{} `json:"data"`
}
