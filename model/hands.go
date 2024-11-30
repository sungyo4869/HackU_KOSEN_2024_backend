package model

type(
	Hand struct {
		UserID int
		CardID int
	}

	ReadHandsResponse struct {
		Hands []Hand
	}
)
