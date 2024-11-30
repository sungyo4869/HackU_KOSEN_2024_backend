package service

import (
	"context"
	"database/sql"

	"github.com/sugyo4869/HackU_KOSEN_2024/model"
)

type HandsService struct {
	db *sql.DB
}

func (s *HandsService)ReadHands(ctx context.Context, userId int) *model.ReadHandsResponse{
	return &model.ReadHandsResponse{}
}
