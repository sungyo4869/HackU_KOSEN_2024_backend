package model

type (
	UserSelected struct {
		Id        int
		UserId    int
		CardId    int
		Attribute string
	}

	UserSelectCardResponse struct {
		Id        int    `json:"Id"`
		CardId    int    `json:"cardId"`
		Attribute string `json:"attribute"`
	}

	ReadUserSelectCardsResponse struct {
		SelectCards []UserSelectCardResponse `json:"selectedCards"`
	}

	UpdateUserSelectCards struct {
		UserSelectCardId int `json:"userSelectedCardId"`
		CardId           int `json:"cardId"`
	}

	UpdateUserSelectCardsRequest struct {
		SelectCards []UpdateUserSelectCards `json:"selectedCards"`
	}
)
