package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/sugyo4869/HackU_KOSEN_2024/model"
	"github.com/sugyo4869/HackU_KOSEN_2024/service"
)

type MatchingHandler struct {
	SCSvc *service.SelectedCardService
	RmSvc *service.RoomService
	ReadyCh chan *websocket.Conn
}

func NewMatchingHandler(SCSvc *service.SelectedCardService, RmSvc *service.RoomService) *MatchingHandler {
	return &MatchingHandler{
		SCSvc: SCSvc,
		RmSvc: RmSvc,
		ReadyCh: make(chan *websocket.Conn, 2),
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // オリジン全部許可
	},
}

func (h *MatchingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade to WebSocket:", err)
		return
	}

	h.ReadyCh <- conn
	defer conn.Close()

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
				log.Println("WebSocket connection closed by client")
				break
			}
			log.Println("Failed to receive message:", err)
			break
		}
	}
}

func (h *MatchingHandler) makeRes(userId []int) (*model.MatchingWSResponse, error) {

	// どっかでユーザーネーム取得しないとなー
	var send model.MatchingWSResponse
	for _, v := range userId {
		cards, err := h.SCSvc.ReadSelectedCard(v)
		if err != nil {
			return nil, err
		}

		var player model.Player
		player.SelectedCards = cards
		player.Username = "user1"

		send.Players = append(send.Players, player)
	}

	room, err := h.RmSvc.CreateRoom(userId)
	if err != nil {
		return nil, err
	}

	send.RoomId = room.RoomId

	return &send, nil
}

func(h *MatchingHandler) StartListening() {

	conn1 := <-h.ReadyCh
	conn2 := <-h.ReadyCh

	log.Println("ふたりそろったよ")

	res, err := h.makeRes([]int{1, 2})
	if err != nil {
		log.Println("Failed to make response:", err)
		return
	}
	
	err = conn1.WriteJSON(res)
	if err != nil {
		log.Println("Failed to send message:", err)
	}

	err = conn2.WriteJSON(res)
	if err != nil {
		log.Println("Failed to send message:", err)
	}
}

