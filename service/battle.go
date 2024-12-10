package service

import (
	"database/sql"
	"fmt"

	"github.com/sugyo4869/HackU_KOSEN_2024/model"
)

type BattleService struct {
	db *sql.DB
}

func NewBattleService(db *sql.DB) *BattleService {
	return &BattleService{
		db: db,
	}
}

func (s *BattleService) InitializeBattle(data *model.InitializeBattleRequest) error {
	const (
		insert = `INSERT INTO battles (
			user_id, room_id, fire_card_id, water_card_id, grass_card_id, 
			kamekame_card_id, nankuru_card_id, random_card_id, hp, result
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`
	)

	_, err := s.db.Exec(insert, data.UserId, data.RoomId, data.FireCardId, data.WaterCardId, data.GrassCardId, data.KameKameCardId, data.NankuruCardId, data.RandomCardId, 3, "pending")
	if err != nil {
		return err
	}

	return nil
}

func (s *BattleService) UpdateBattle(userId int, roomId int, column string, hp int) error {
	query := fmt.Sprintf("UPDATE battles SET %s = NULL hp = ? WHERE user_id = ? AND battle_id = ?;", column)

	_, err := s.db.Exec(query, hp, userId, roomId)
	if err != nil {
		return err
	}

	return nil
}

func (s *BattleService) ReadBattle(userId int, roomId int) (*model.Battle, error) {
	const read = `select * from cards where user_id = ? and room_id = ?`

	var battle model.Battle

	row := s.db.QueryRow(read, userId,roomId) 
	if err := row.Scan(&battle.Battle_id, &battle.UserId, &battle.RandomId, &battle.FireCardId, &battle.WaterCardId, &battle.GrassCardId, &battle.KameKameCardId, &battle.NankuruCardId, &battle.RandomId, &battle.Hp, &battle.Result); err != nil {
		return nil, err
	}

	return &battle, nil

}
