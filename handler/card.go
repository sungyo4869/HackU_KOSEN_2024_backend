package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/sugyo4869/HackU_KOSEN_2024/model"
	"github.com/sugyo4869/HackU_KOSEN_2024/service"
)

type CardHandler struct {
	svc *service.CardService
}

func NewCardHandler(svc service.CardService) *CardHandler {
	return &CardHandler{
		svc: &svc,
	}
}

func (h *CardHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	type userIDKey struct{}

	switch r.Method {
	case http.MethodGet:
		var res model.ReadCardsResponse
		userID, ok := r.Context().Value(userIDKey{}).(int64)
		if !ok {
			http.Error(w, "ユーザーIDが見つかりません", http.StatusUnauthorized)
			return
		}

		cards, err := h.svc.ReadCard(r.Context(), int(userID), []int{})
		if err != nil {
			http.Error(w, "cards is not found", http.StatusNotFound)
			log.Println("cards is not found, err = ", err)
		}

		res.Cards = *cards
		err = json.NewEncoder(w).Encode(&res)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Println("Login: Filed to encoding json, err = ", err)
			return
		}

	}
}
