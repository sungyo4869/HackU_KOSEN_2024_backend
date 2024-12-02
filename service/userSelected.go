package service

import (
	"context"
	"database/sql"

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

func (s *UserSelectedService) ReadUserSelected(ctx context.Context, userId int) (*[]model.UserSelected, error) {
	query := `select * from user_selected where user_id = ?`

	rows, err := s.db.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var selections []model.UserSelected
	for rows.Next() {
		var selection model.UserSelected
		if err := rows.Scan(&selection.Id, &selection.UserId, &selection.CardId, &selection.Attribute); err != nil {
			return nil, err
		}

		selections = append(selections, selection)
	}

	return &selections, nil
}
