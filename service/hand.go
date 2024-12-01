package service

import (
	"context"
	"database/sql"

	"github.com/sugyo4869/HackU_KOSEN_2024/model"
)

type HandService struct {
	db *sql.DB
}

func NewHandService(db *sql.DB) *HandService {
	return &HandService{
		db: db,
	}
}

func (s *HandService) ReadHand(ctx context.Context, userId int) (*[]model.Hand, error) {
	query := `select * from hands where user_id = ?`

	rows, err := s.db.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hands []model.Hand
	for rows.Next() {
		var hand model.Hand
		if err := rows.Scan(&hand.ID, &hand.UserID, &hand.CardID); err != nil {
			return nil, err
		}

		hands = append(hands, hand)
	}

	return &hands, nil
}
