package model

type (
	User struct {
		UserId int    `json:"user-id"`
		Name   string `json:"username"`
	}

	ReadUserRequest struct {
		Name string
	}

	ReadUserResponse struct {
		UserId int `json:"user-id"`
	}
)
