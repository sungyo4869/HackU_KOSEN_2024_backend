package model

type (
	Card struct {
		Id      int64
		UserId  int64
		Picture string
		Name    string
	}

	ReadCardsResponse struct {
		Cards []Card
	}
)
