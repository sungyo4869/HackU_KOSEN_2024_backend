package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sugyo4869/HackU_KOSEN_2024/model"
)

type HealthzHandler struct{}

func NewHealthzHandler() *HealthzHandler {
	return &HealthzHandler{}
}

func (h *HealthzHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res := &model.HealthzResponse{
		Message: "OK",
	}

	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
}
