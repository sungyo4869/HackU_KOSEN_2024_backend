package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/sugyo4869/HackU_KOSEN_2024/handler/middleware"
	"github.com/sugyo4869/HackU_KOSEN_2024/model"
	"github.com/sugyo4869/HackU_KOSEN_2024/service"
)

type UserSelectedHandler struct {
	svc *service.UserSelectedService
}

func NewUserSelectedHandler(svc service.UserSelectedService) *UserSelectedHandler {
	return &UserSelectedHandler{
		svc: &svc,
	}
}

func (h *UserSelectedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		var res model.ReadUserSelectedCardsResponse

		userId, ok := r.Context().Value(middleware.UserIDKey{}).(int64)
		if !ok {
			http.Error(w, "ユーザーIDが見つかりません", http.StatusUnauthorized)
			return
		}

		selections, err := h.svc.ReadUserSelected(r.Context(), int(userId))
		if err != nil {
			http.Error(w, "selections is not found", http.StatusNotFound)
			log.Println("selections is not found, err = ", err)
		}

		// res.SelectedCards = *selections
		res.SelectedCards = *selections
		err = json.NewEncoder(w).Encode(&res)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Println("Login: Filed to encoding json, err = ", err)
			return
		}
	}
}
