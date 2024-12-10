package service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sugyo4869/HackU_KOSEN_2024/model"
)

type UserSelectService struct {
	db *sql.DB
}

func NewUserSelectService(db *sql.DB) *UserSelectService {
	return &UserSelectService{
		db: db,
	}
}

func (s *UserSelectService) ReadUserSelect(ctx context.Context, userId int) (*[]model.UserSelectCardResponse, error) {
	query := `SELECT id, card_id, attribute from user_selected where user_id = ?`

	rows, err := s.db.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var selectCards []model.UserSelectCardResponse
	for rows.Next() {
		var selectCard model.UserSelectCardResponse
		if err := rows.Scan(&selectCard.SelectCardId, &selectCard.CardId, &selectCard.Attribute); err != nil {
			return nil, err
		}

		selectCards = append(selectCards, selectCard)
	}

	return &selectCards, nil
}

func (s *UserSelectService) UpdateUserSelect(ctx context.Context, userId int, userSelectCards []model.UpdateUserSelectCards) (*[]model.UserSelectCardResponse, error) {
	checkOwnership := `SELECT COUNT(*) FROM cards WHERE id = ? AND user_id = ?`
	update := `UPDATE user_selected SET card_id = ? WHERE id = ? AND user_id = ?`
	confirm := `SELECT id, card_id, attribute FROM user_selected WHERE user_id = ?`

	var resCards []model.UserSelectCardResponse

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	for _, userSelectCard := range userSelectCards {
		var count int
		err := tx.QueryRowContext(ctx, checkOwnership, userSelectCard.CardId, userId).Scan(&count)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		if count == 0 {
			tx.Rollback()
			return nil, fmt.Errorf("card_id %d does not belong to user_id %d", userSelectCard.CardId, userId)
		}
		_, err = tx.ExecContext(ctx, update, userSelectCard.CardId, userSelectCard.SelectCardId, userId)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	rows, err := s.db.QueryContext(ctx, confirm, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var resCard model.UserSelectCardResponse
		err := rows.Scan(
			&resCard.SelectCardId,
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
