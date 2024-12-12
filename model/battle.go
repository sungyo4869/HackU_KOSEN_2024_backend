package model

import "database/sql"

type (
	Battle struct {
		Battle_id      int64
		UserId         int64
		RoomId         int64
		RedCardId      sql.NullInt64
		BlueCardId     sql.NullInt64
		GreenCardId    sql.NullInt64
		KameKameCardId sql.NullInt64
		NankuruCardId  sql.NullInt64
		RandomCardId       sql.NullInt64
		RandomAttribute string
		Hp             int
		Result         string
	}

	InitializeBattleRequest struct {
		UserId         int64
		RoomId         int64
		RedCardId      sql.NullInt64
		BlueCardId     sql.NullInt64
		GreenCardId    sql.NullInt64
		KameKameCardId sql.NullInt64
		NankuruCardId  sql.NullInt64
		RandomCardId   sql.NullInt64
		RandomAttribute string
		Result         string
	}
)
