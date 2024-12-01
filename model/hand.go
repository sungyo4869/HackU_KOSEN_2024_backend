package model

type (
	Hand struct {
		ID     int
		UserID int
		CardID int
	}

	ReadHandResponse struct {
		Hands []Hand
	}
)
