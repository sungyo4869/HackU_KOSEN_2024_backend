package model

type (
	UserSelected struct {
		Id        int
		UserId    int
		CardId    int
		Attribute string
	}

	UserSelectCardResponse struct {
		SelectCardId int    `json:"select-cardId"`
		CardId       int    `json:"card-id"`
		Attribute    string `json:"attribute"`
	}

	ReadUserSelectCardsResponse struct {
		SelectCards []UserSelectCardResponse `json:"select-cards"`
	}

	UpdateUserSelectCards struct {
		Attribute string `json:"attribute"`
		CardId    int    `json:"card-id"`
	}

	UpdateUserSelectCardsRequest struct {
		SelectCards []UpdateUserSelectCards `json:"select-cards"`
	}
)
