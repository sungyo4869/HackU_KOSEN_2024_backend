package service

import (
	"context"
	"database/sql"

	"github.com/sugyo4869/HackU_KOSEN_2024/model"
)

type UserService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		db: db,
	}
}

func (s *UserService) ReadUser(ctx context.Context, username string) (*model.User, error) {
	const query = `select * from users where username = ?`
	var user model.User

	row := s.db.QueryRowContext(ctx, query, username)
	if err := row.Scan(&user.UserId, &user.Name, &user.Password); err != nil {
		return nil, err
	}

	return &user, nil

}
