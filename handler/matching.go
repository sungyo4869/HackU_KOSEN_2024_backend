package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/sugyo4869/HackU_KOSEN_2024/model"
	"github.com/sugyo4869/HackU_KOSEN_2024/service"
)

type MatchingHandler struct {
	SCSvc *service.SelectedCardService
	RmSvc *service.RoomService
}

func NewMatchingHandler(SCSvc *service.SelectedCardService, RmSvc *service.RoomService) *MatchingHandler {
	return &MatchingHandler{
		SCSvc: SCSvc,
		RmSvc: RmSvc,
	}
}

func (h *MatchingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var res model.MatchingWSResponse

	cards, err := h.SCSvc.ReadSelectedCard(r.Context(), 1)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		log.Println("matching: err = ", err)
		return
	}

	var player model.Player

	player.SelectedCards = cards
	player.Username = "user1"
	res.Players = append(res.Players, player)

	cards, err = h.SCSvc.ReadSelectedCard(r.Context(), 1)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		log.Println("matching: err = ", err)
		return
	}

	player.SelectedCards = cards
	player.Username = "user1"

	res.Players = append(res.Players, player)

	room, err := h.RmSvc.CreateRoom(r.Context(), []int{1,2})
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		log.Println("matching: err = ", err)
		return
	}

	res.RoomId = room.RoomId
	
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

}
