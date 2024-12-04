package service

import (
	"context"
	"database/sql"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
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

func (s *UserService) ReadUser(ctx context.Context, username string, password string) (*model.LoginResponse, error) {
	const query = `select * from users where username = ? AND password = ?`
	var user model.User

	row := s.db.QueryRowContext(ctx, query, username, password)
	if err := row.Scan(&user.UserId, &user.Name, &user.Password); err != nil {
		return nil, err
	}

	var res model.LoginResponse

	// JWTトークン生成
	token, err := s.CreateToken(user.UserId)
	if err != nil {
		return nil, err
	}
	res.Token = token

	return &res, nil
}

func (s *UserService) ReadUserWithId(userId int) (*model.ReadUserWithIdResponse, error){
	const query = `select username from users where id = ?`
	var user model.ReadUserWithIdResponse

	row := s.db.QueryRow(query, userId)
	if err := row.Scan(&user.Username); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UserService) CreateToken(userId int) (string, error) {
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	token.Claims = jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(72 * time.Hour).Unix(),
	}

	var secretKey = os.Getenv("SECRET_KEY")
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
