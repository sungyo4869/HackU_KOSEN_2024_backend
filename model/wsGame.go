package model

import "database/sql"

type (
	GameWSRequest struct {
		RoomId    int64  `json:"room-id"`
		UserId    int64  `json:"user-id"`
		Attribute string `json:"attribute"`
		CardId    int64  `json:"card-id"`
	}

	ShogunRequest struct {
		RoomId   int64 `json:"room-id"`
		UserId   int64 `json:"user-id"`
		ShogunId int64 `json:"shogun-id"`
	}

	Shogun struct {
		RoomId   int64 `json:"room-id"`
		UserId   int64 `json:"user-id"`
		ShogunId int64 `json:"shogun-id"`
	}

	ShogunResponse struct {
		Players []Shogun `json:"players"`
	}

	GameWSResponse struct {
		Results []GameResult `json:"results"`
	}

	GameResult struct {
		UserId          int64         `json:"user-id"`
		Hp              int           `json:"hp"`
		SelectAttribute string        `json:"select-attribute"`
		SelectCardId    int64         `json:"select-card-id"`
		TurnResult      string        `json:"turn-result"`
		RedCardId       sql.NullInt64 `json:"red-card-id"`
		BlueCardId      sql.NullInt64 `json:"blue-card-id"`
		GreenCardId     sql.NullInt64 `json:"green-card-id"`
		KameKameCardId  sql.NullInt64 `json:"kamekame-card-id"`
		NankuruCardId   sql.NullInt64 `json:"nankuru-card-id"`
		RandomCardId    sql.NullInt64 `json:"random-card-id"`
	}
)
