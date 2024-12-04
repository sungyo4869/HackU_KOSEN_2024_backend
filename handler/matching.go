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
	UsrSvc *service.UserService
	ReadyCh chan *websocket.Conn
	Player chan *int
}

func NewMatchingHandler(SCSvc *service.SelectedCardService, RmSvc *service.RoomService, UsrSvc service.UserService) *MatchingHandler {
	return &MatchingHandler{
		SCSvc: SCSvc,
		RmSvc: RmSvc,
		UsrSvc: &UsrSvc,
		ReadyCh: make(chan *websocket.Conn, 2),
		Player: make(chan *int, 2),
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

	defer conn.Close()

	msg := model.ReadUserResponse{}
	err = conn.ReadJSON(&msg)
	if err != nil{
		log.Println("Failed to receive json:", err)
		return
	}

	h.Player <- &msg.UserId
	h.ReadyCh <- conn

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

	// user-idからusernameを取得するサービスを実装する
	
	var send model.MatchingWSResponse
	for _, v := range userId {
		user, err := h.UsrSvc.ReadUserWithId(v)
		if err != nil {
			return nil, err
		}
		cards, err := h.SCSvc.ReadSelectedCard(v)
		if err != nil {
			return nil, err
		}

		var player model.Player
		player.SelectedCards = cards
		player.Username = user.Username

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
	for {
		conn1 := <-h.ReadyCh
		conn2 := <-h.ReadyCh

		player1 := <-h.Player
		player2 := <-h.Player

		res, err := h.makeRes([]int{*player1, *player2})
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

		err = conn1.Close()
		if err != nil {
			log.Println("Failed to close conn1:", err)
		}

		err = conn2.Close()
		if err != nil {
			log.Println("Failed to close conn2:", err)
		}
	}
}

