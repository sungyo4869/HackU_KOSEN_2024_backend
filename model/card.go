package model

type (
	Card struct {
		Id      int64  `json:"card-id"`
		UserId  int64  `json:"user-id"`
		Picture string `json:"picture"`
		Name    string `json:"name"`
	}

	ReadCardsResponse struct {
		Cards []Card
	}
)
