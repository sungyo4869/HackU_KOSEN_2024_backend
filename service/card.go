package service

import (
	"context"
	"database/sql"

	"github.com/sugyo4869/HackU_KOSEN_2024/model"
)

type CardService struct {
	db *sql.DB
} 

func NewCardService(db *sql.DB) *CardService {
	return &CardService{
		db: db,
	}
}

func(s *CardService) ReadCard(ctx context.Context, userId int) (*[]model.Card, error){
	query := `select * from cards where user_id = ?`

    rows, err := s.db.QueryContext(ctx, query, userId)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var cards []model.Card
    for rows.Next() {
        var card model.Card
        if err := rows.Scan(&card.Id, &card.UserId, &card.Name, &card.Picture); err != nil {
            return nil, err
        }
        cards = append(cards, card)
    }

	return &cards, nil	
}
