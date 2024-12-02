package model

type (
	User struct {
		UserId   int    `json:"user-id"`
		Name     string `json:"username"`
		Password string `json:"password"`
	}

	ReadUserResponse struct {
		UserId int `json:"user-id"`
	}
)
