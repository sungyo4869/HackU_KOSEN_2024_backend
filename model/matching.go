package model

type MatchingWSResponse struct {
	RoomId  int      `json:"room-id"`
	Players []Player `json:"players"`
}

type Player struct {
	Username      string         `json:"username`
	SelectedCards []SelectedCard `json:"selected-cards"`
}
