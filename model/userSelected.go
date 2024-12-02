package model

type (
	UserSelected struct {
		Id        int
		UserId    int
		CardId    int
		Attribute string
	}

	ReadHandsResponse struct {
		SelectedCards []UserSelected `json:"selected-cards"`
	}
)
