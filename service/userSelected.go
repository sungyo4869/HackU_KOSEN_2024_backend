package service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sugyo4869/HackU_KOSEN_2024/model"
)

type UserSelectedService struct {
	db *sql.DB
}

func NewUserSelectedService(db *sql.DB) *UserSelectedService {
	return &UserSelectedService{
		db: db,
	}
}

func (s *UserSelectedService) ReadUserSelected(ctx context.Context, userId int) (*[]model.UserSelectedCardResponse, error) {
	query := `SELECT id, card_id, attribute from user_selected where user_id = ?`

	rows, err := s.db.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var selections []model.UserSelectedCardResponse
	for rows.Next() {
		var selection model.UserSelectedCardResponse
		if err := rows.Scan(&selection.Id, &selection.CardId, &selection.Attribute); err != nil {
			return nil, err
		}

		selections = append(selections, selection)
	}

	return &selections, nil
}

func (s *UserSelectedService) UpdateUserSelected(ctx context.Context, userId int, userSelectedCards []model.UpdateUserSelectedCards) (*[]model.UserSelectedCardResponse, error) {
	checkOwnership := `SELECT COUNT(*) FROM cards WHERE id = ? AND user_id = ?`
	update := `UPDATE user_selected SET card_id = ? WHERE id = ? AND user_id = ?`
	confirm := `SELECT id, card_id, attribute FROM user_selected WHERE user_id = ?`

	var resCards []model.UserSelectedCardResponse

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	for _, userSelectedCard := range userSelectedCards {
		var count int
		err := tx.QueryRowContext(ctx, checkOwnership, userSelectedCard.CardId, userId).Scan(&count)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		if count == 0 {
			tx.Rollback()
			return nil, fmt.Errorf("card_id %d does not belong to user_id %d", userSelectedCard.CardId, userId)
		}
		_, err = tx.ExecContext(ctx, update, userSelectedCard.CardId, userSelectedCard.UserSelectedCardId, userId)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	// id, err := result.LastInsertId()
	// if err != nil {
	// 	return nil, err
	// }

	rows, err := s.db.QueryContext(ctx, confirm, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var resCard model.UserSelectedCardResponse
		err := rows.Scan(
			&resCard.Id,
			&resCard.CardId,
			&resCard.Attribute,
		)
		if err != nil {
			return nil, err
		}
		resCards = append(resCards, resCard)
	}

	return &resCards, nil
}
