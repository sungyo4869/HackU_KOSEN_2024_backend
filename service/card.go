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

func(s *CardService) ReadCard(ctx context.Context, userId int, cardId []int) (*[]model.Card, error){
	const (
        read = `select * from cards where user_id = ?`
        readWithId = `select * from cards where user_id = ? and card_id in (?, ?, ?, ?, ?, ?)`
    )

    var rows *sql.Rows
    var err error

    if len(cardId) == 0{
        rows, err = s.db.QueryContext(ctx, read, userId)
    } else {
        rows, err = s.db.QueryContext(ctx, readWithId, cardId[0], cardId[1], cardId[2], cardId[3], cardId[4], cardId[5])
    }

    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var cards []model.Card
    for rows.Next() {
        var card model.Card
        if err := rows.Scan(&card.Id, &card.UserId, &card.Picture, &card.Name); err != nil {
            return nil, err
        }
        cards = append(cards, card)
    }

	return &cards, nil	
}
