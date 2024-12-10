package model

type (
	UserSelected struct {
		Id        int
		UserId    int
		CardId    int
		Attribute string
	}

	UserSelectCardResponse struct {
		SelectCardId int    `json:"selectCardId"`
		CardId       int    `json:"cardId"`
		Attribute    string `json:"attribute"`
	}

	ReadUserSelectCardsResponse struct {
		SelectCards []UserSelectCardResponse `json:"selectedCards"`
	}

	UpdateUserSelectCards struct {
		SelectCardId int `json:"selectedCardId"`
		CardId       int `json:"cardId"`
	}

	UpdateUserSelectCardsRequest struct {
		SelectCards []UpdateUserSelectCards `json:"selectedCards"`
	}
)
