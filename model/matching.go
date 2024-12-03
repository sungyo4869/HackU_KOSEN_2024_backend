package model

type MatchingWSResponse struct {
	RoomId  int `json:"room-id"`
	Players []Player `json:"players"`
}

type Player struct {
	User        User `json:"user"`
	SelectedCards []SelectedCard `json:"selected-cards"`
}
