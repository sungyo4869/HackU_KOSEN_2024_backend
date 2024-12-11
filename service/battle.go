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
			user_id, room_id, red_card_id, blue_card_id, green_card_id, 
			kamekame_card_id, nankuru_card_id, random_card_id, hp, result
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`
	)

	_, err := s.db.Exec(insert, data.UserId, data.RoomId, data.RedCardId, data.BlueCardId, data.GreenCardId, data.KameKameCardId, data.NankuruCardId, data.RandomCardId, 3, "pending")
	if err != nil {
		return err
	}

	return nil
}

func (s *BattleService) UpdateBattle(userId int64, roomId int64, attribute string, hp int) (*model.Battle, error) {
	query := fmt.Sprintf("UPDATE battles SET %s = NULL, hp = ? WHERE user_id = ? AND room_id = ?;", attribute+"_card_id")
	const confirm = `SELECT * FROM battles WHERE user_id = ? AND room_id = ?`

	// UPDATE クエリの実行
	result, err := s.db.Exec(query, hp, userId, roomId)
	if err != nil {
		return nil, err
	}

	// 変更された行数を確認
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, fmt.Errorf("no rows affected")
	}

	var res model.Battle
	row := s.db.QueryRow(confirm, userId, roomId)
	if err := row.Scan(&res.Battle_id, &res.UserId, &res.RoomId, &res.RedCardId, &res.BlueCardId, &res.GreenCardId, &res.KameKameCardId, &res.NankuruCardId, &res.RandomCardId, &res.Hp, &res.Result); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no battle found for user_id %d and room_id %d", userId, roomId)
		}
		return nil, err
	}

	return &res, nil
}

func (s *BattleService) UpdateResult(userId int64, roomId int64, result string) (*model.Battle, error) {
	const (
		query   = `UPDATE battles SET result = ? WHERE user_id = ? AND room_id = ?`
		confirm = `SELECT * FROM battles WHERE user_id = ? AND room_id = ?`
	)

	execResult, err := s.db.Exec(query, result, userId, roomId)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := execResult.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, fmt.Errorf("no rows affected")
	}

	var res model.Battle
	row := s.db.QueryRow(confirm, userId, roomId)
	if err := row.Scan(&res.Battle_id, &res.UserId, &res.RoomId, &res.RedCardId, &res.BlueCardId, &res.GreenCardId, &res.KameKameCardId, &res.NankuruCardId, &res.RandomCardId, &res.Hp, &res.Result); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no battle found for user_id %d and room_id %d", userId, roomId)
		}
		return nil, err
	}

	return &res, nil
}


func (s *BattleService) ReadBattle(userId int64, roomId int64) (*model.Battle, error) {
	const read = `select * from battles where user_id = ? and room_id = ?`

	var battle model.Battle

	row := s.db.QueryRow(read, userId, roomId)
	if err := row.Scan(&battle.Battle_id, &battle.UserId, &battle.RoomId, &battle.RedCardId, &battle.BlueCardId, &battle.GreenCardId, &battle.KameKameCardId, &battle.NankuruCardId, &battle.RandomCardId, &battle.Hp, &battle.Result); err != nil {
		return nil, err
	}

	return &battle, nil

}
