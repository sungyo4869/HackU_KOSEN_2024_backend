package model

type (
	MatchingUserRequest struct {
		UserId int64 `json:"user-id"`
	}

	MatchingWSResponse struct {
		RoomId  int64    `json:"room-id"`
		Players []Player `json:"players"`
	}

	Player struct {
		Username      string         `json:"username"`
		SelectedCards []SelectedCard `json:"selected-cards"`
	}
)
