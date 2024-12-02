package model

type(
	SelectedCard struct {
		UserID int 
		CardID int 
		Attribute string 
	}

	ReadHandsResponse struct {
		SelectedCards []SelectedCard `json:"selected-cards"`
	}
)
