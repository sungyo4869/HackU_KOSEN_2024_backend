package model

type (
	Battle struct {
		Battle_id      int
		UserId         int
		RoomId         int
		FireCardId     int
		WaterCardId    int
		GrassCardId    int
		KameKameCardId int
		NankuruCardId  int
		RandomId       int
		Hp             int
		Result         string
	}

	InitializeBattleRequest struct {
		UserId         int
		RoomId         int
		FireCardId     int
		WaterCardId    int
		GrassCardId    int
		KameKameCardId int
		NankuruCardId  int
		RandomCardId   int
		Result         string
	}
)
