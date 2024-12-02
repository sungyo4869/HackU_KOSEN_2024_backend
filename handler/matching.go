package handler

import (
	"net/http"

	"github.com/sugyo4869/HackU_KOSEN_2024/model"
)

type MatchingHandler struct{}

func NewMatchingHandler() *MatchingHandler {
	return &MatchingHandler{}
}

func (h *MatchingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) *model.MatchingWSResponse {

	return &model.MatchingWSResponse{}
}
