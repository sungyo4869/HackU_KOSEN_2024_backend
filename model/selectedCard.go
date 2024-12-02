package model

type SelectedCard struct {
	Id        int    `json:"active-card-id"`
	Attribute string `json:"attribute"`
	Name      string `json:"name"`
	Picture   string `json:"picture"`
}
