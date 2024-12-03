package service


import (
	"context"
	"database/sql"

	"github.com/sugyo4869/HackU_KOSEN_2024/model"
)

type RoomService struct {
	db *sql.DB
}

func NewRoomService(db *sql.DB) *RoomService {
	return &RoomService{
		db: db,
	}
}

func (s *RoomService) CreateRoom(ctx context.Context, userid []int) (*model.Room, error) {
	const (
		insert  = `INSERT INTO rooms(user1_id, user2_id) VALUES(?, ?)`
		confirm = `SELECT * FROM rooms WHERE room_id = ?`
	)

	result, err := s.db.ExecContext(ctx, insert, userid[0], userid[1])
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	var room model.Room
	row := s.db.QueryRowContext(ctx, confirm, id)

	err = row.Scan(
		&room.RoomId,
		&room.UserId1,
		&room.UserId2,
	)
	
	if err != nil {
		return nil, err
	}

	return &room, nil
}
