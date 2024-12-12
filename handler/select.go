package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/sugyo4869/HackU_KOSEN_2024/handler/middleware"
	"github.com/sugyo4869/HackU_KOSEN_2024/model"
	"github.com/sugyo4869/HackU_KOSEN_2024/service"
)

type SelectHandler struct {
	svc *service.UserSelectService
}

func NewSelectHandler(svc service.UserSelectService) *SelectHandler {
	return &SelectHandler{
		svc: &svc,
	}
}

func (h *SelectHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		var res model.ReadUserSelectCardsResponse

		userId, ok := r.Context().Value(middleware.UserIDKey{}).(int64)
		if !ok {
			http.Error(w, "ユーザーIDが見つかりません", http.StatusUnauthorized)
			return
		}

		selected, err := h.svc.ReadUserSelect(r.Context(), int(userId))
		if err != nil {
			http.Error(w, "selected is not found", http.StatusNotFound)
			log.Println("selected is not found, err = ", err)
		}

		res.SelectCards = *selected
		err = json.NewEncoder(w).Encode(&res)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Println("Login: Filed to encoding json, err = ", err)
			return
		}

	case http.MethodPut:
		var req model.UpdateUserSelectCardsRequest
		var res model.ReadUserSelectCardsResponse

		userId, ok := r.Context().Value(middleware.UserIDKey{}).(int64)
		if !ok {
			http.Error(w, "ユーザーIDが見つかりません", http.StatusUnauthorized)
			return
		}

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Bad Request: Invalid JSON", http.StatusBadRequest)
			log.Println("Failed to decode request body, err =", err)
			return
		}

		selected, err := h.svc.UpdateUserSelect(r.Context(), int(userId), req.SelectCards)
		if err != nil {
			http.Error(w, "selected is not found", http.StatusNotFound)
			log.Println("selected is not found, err = ", err)
			return
		}

		res.SelectCards = *selected
		err = json.NewEncoder(w).Encode(&res)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Println("Login: Filed to encoding json, err = ", err)
			return
		}
	}
}
