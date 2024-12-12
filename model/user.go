package model

type (
	User struct {
		UserId   int64  `json:"user-id"`
		Name     string `json:"username"`
		Password string `json:"password"`
	}

	ReadUserResponse struct {
		UserId int64 `json:"user-id"`
	}

	ReadUserWithIdResponse struct {
		Username string `json:"username"`
	}

	LoginResponse struct {
		UserId int64  `json:"user-id"`
		Token  string `json:"token"`
	}
)
