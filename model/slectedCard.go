package model

type (
	SelectedCard struct {
		ID        int
		UserID    int
		CardID    int
		Attribute string
	}

	ReadHandsResponse struct {
		SelectedCards []SelectedCard `json:"selected-cards"`
	}
)
