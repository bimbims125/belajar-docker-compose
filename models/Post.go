package models

type Post struct {
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type SuccessResponseMessage struct {
	Message string `json:"message"`
}

type ErrorResponseMessage struct {
	Message string `json:"message"`
}

type SuccessResponse struct {
	Data interface{} `json:"data"`
}
