package model

type (
	User struct {
		UserId int    `json:"user-id"`
		Name   string `json:"username"`
	}

	ReadUserResponse struct {
		UserId int `json:"user-id"`
	}
)
