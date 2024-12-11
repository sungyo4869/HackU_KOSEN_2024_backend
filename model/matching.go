package model

type MatchingWSResponse struct {
	RoomId  int64      `json:"room-id"`
	Players []Player `json:"players"`
}

type Player struct {
	Username      string         `json:"username"`
	SelectedCards []SelectedCard `json:"selected-cards"`
}
