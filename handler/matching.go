package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sugyo4869/HackU_KOSEN_2024/service"
)

type MatchingHandler struct {
	svc *service.SelectedCardService
}

func NewMatchingHandler(svc *service.SelectedCardService) *MatchingHandler {
	return &MatchingHandler{
		svc: svc,
	}
}

func (h *MatchingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	res, err := h.svc.ReadSelectedCard(r.Context(), 1)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

}
