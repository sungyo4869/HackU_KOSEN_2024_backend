package model

type (
	UserSelected struct {
		Id        int
		UserId    int
		CardId    int
		Attribute string
	}

	UserSelectedCardResponse struct {
		CardId    int    `json:"cardId"`
		Attribute string `json:"attribute"`
	}

	ReadUserSelectedCardsResponse struct {
		SelectedCards []UserSelectedCardResponse `json:"selectedCards"`
	}
)
