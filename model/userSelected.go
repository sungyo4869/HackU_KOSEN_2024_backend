package model

type (
	UserSelected struct {
		Id        int
		UserId    int
		CardId    int
		Attribute string
	}

	UserSelectedCardResponse struct {
		Id        int    `json:"Id"`
		CardId    int    `json:"cardId"`
		Attribute string `json:"attribute"`
	}

	ReadUserSelectedCardsResponse struct {
		SelectedCards []UserSelectedCardResponse `json:"selectedCards"`
	}

	UpdateUserSelectedCards struct {
		UserSelectedCardId int `json:"userSelectedCardId"`
		CardId             int `json:"cardId"`
	}

	UpdateUserSelectedCardsRequest struct {
		SelectedCards []UpdateUserSelectedCards `json:"selectedCards"`
	}
)
