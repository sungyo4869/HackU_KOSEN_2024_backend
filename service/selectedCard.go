package service

import (
	"context"
	"database/sql"

	"github.com/sugyo4869/HackU_KOSEN_2024/model"
)

type SelectedCardService struct {
	db *sql.DB
}

func NewSelectedCardService(db *sql.DB) *SelectedCardService {
	return &SelectedCardService{
		db: db,
	}
}

func (s *SelectedCardService) ReadSelectedCard(ctx context.Context, userId int) ([]model.SelectedCard, error){
	query := `SELECT us.attribute, c.card_name, c.id AS card_id, c.picture FROM user_selected us JOIN cards c ON us.card_id = c.id WHERE us.user_id = ?;`

	rows, err := s.db.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var selectedCards []model.SelectedCard
	for rows.Next() {
		var selectedCard model.SelectedCard
		if err := rows.Scan(&selectedCard.Attribute, &selectedCard.Name, &selectedCard.CardId, &selectedCard.Picture); err != nil {
			return nil, err
		}

		selectedCards = append(selectedCards, selectedCard)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	
	return selectedCards, nil
}

