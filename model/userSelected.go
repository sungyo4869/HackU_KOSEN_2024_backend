package model

type (
	UserSelected struct {
		ID        int
		UserID    int
		CardID    int
		Attribute string
	}

	ReadHandsResponse struct {
		SelectedCards []UserSelected `json:"selected-cards"`
	}
)
