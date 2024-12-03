package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

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
	switch r.Method {
	case http.MethodGet:
		var res model.ReadCardsResponse
		params := r.URL.Query()
		id := params.Get("user-id")
		if id == "" {
			http.Error(w, "user-id is missing", http.StatusBadRequest)
			log.Println("Failed to get parameters")
		}

		intId, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "user-id is missing", http.StatusBadRequest)
			log.Println("Failed to get parameters")
		}

		cards, err := h.svc.ReadCard(r.Context(), intId, []int{})
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
