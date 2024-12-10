package model

type (
	GameRequest struct {
		RoomId  int    `json:"room-id"`
		UserId  int    `json:"user-id"`
		PutCard string `json:"put-card"`
	}
	
	// 以下まだ考えてない
	GameResponse struct {
		Message string
	}
)
