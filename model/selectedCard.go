package model

type SelectedCard struct {
	CardId    int    `json:"card-id"`
	Attribute string `json:"attribute"`
	Name      string `json:"name"`
	Picture   string `json:"picture"`
}
