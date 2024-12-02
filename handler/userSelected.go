package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/sugyo4869/HackU_KOSEN_2024/model"
	"github.com/sugyo4869/HackU_KOSEN_2024/service"
)

type HandHandler struct {
	svc *service.HandService
}

func NewHandHandler(svc service.HandService) *HandHandler {
	return &HandHandler{
		svc: &svc,
	}
}

func (h *HandHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		var res model.ReadHandsResponse
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

		selections, err := h.svc.ReadHand(r.Context(), intId)
		if err != nil {
			http.Error(w, "selections is not found", http.StatusNotFound)
			log.Println("selections is not found, err = ", err)
		}

		res.SelectedCards = *selections
		err = json.NewEncoder(w).Encode(&res)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Println("Login: Filed to encoding json, err = ", err)
			return
		}
	}
}
