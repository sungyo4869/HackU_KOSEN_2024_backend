package model

import "database/sql"

type (
	GameRequest struct {
		RoomId    int64  `json:"room-id"`
		UserId    int64  `json:"user-id"`
		Attribute string `json:"attribute"`
	}

	GameResponse struct {
		Results []GameResult
	}

	GameResult struct {
		UserId          int64         `json:"user-id"`
		Hp              int           `json:"hp"`
		SelectAttribute string        `json:"select-attribute"`
		TurnResult      string        `json:"turn-result"`
		RedCardId       sql.NullInt64 `json:"red-card-id"`
		BlueCardId      sql.NullInt64 `json:"blue-card-id"`
		GreenCardId     sql.NullInt64 `json:"green-card-id"`
		KameKameCardId  sql.NullInt64 `json:"kamekame-card-id"`
		NankuruCardId   sql.NullInt64 `json:"nankuru-card-id"`
		RandomCardId    sql.NullInt64 `json:"random-card-id"`
	}
)
