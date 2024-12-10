package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/sugyo4869/HackU_KOSEN_2024/model"
	"github.com/sugyo4869/HackU_KOSEN_2024/service"
)

type GameHandler struct {
	SCSvc   *service.SelectedCardService
	RmSvc   *service.RoomService
	UsrSvc  *service.UserService
	ReadyCh chan *websocket.Conn
	UserId  chan *int
	RoomId  chan *int
	PutCard chan *string
}

type player struct {
	conn    *websocket.Conn
	UserId  *int
	RoomId  *int
	PutCard *string
}

func NewGameHandler(SCSvc service.SelectedCardService, RmSvc service.RoomService, UsrSvc service.UserService) *GameHandler {
	return &GameHandler{
		RmSvc:   &RmSvc,
		UsrSvc:  &UsrSvc,
		ReadyCh: make(chan *websocket.Conn, 2),
		UserId:  make(chan *int, 2),
		RoomId:  make(chan *int, 2),
		PutCard: make(chan *string, 2),
	}
}

func (h *GameHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade to WebSocket:", err)
		return
	}

	defer conn.Close()

	for {
		var msg model.GameRequest

		err = conn.ReadJSON(&msg)

		h.UserId <- &msg.UserId
		h.RoomId <- &msg.RoomId
		h.PutCard <- &msg.PutCard
		h.ReadyCh <- conn

		log.Println("おくりだしました")

		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
				log.Println("WebSocket connection closed by client")
				break
			}

			log.Println("Failed to receive json:", err)
			return
		}
	}
}

func (h *GameHandler) makeRes(player1 *player, player2 *player) *model.GameResponse {
	// ReadBattleとかもつかって、ちゃんとレスポンスかえすようにしないとな
	log.Println(*player1.PutCard, *player1.UserId)
	log.Println(*player2.PutCard, *player2.UserId)

	if *player1.PutCard == *player2.PutCard {
		return &model.GameResponse{
			Message: "draw",
		}
	} else {
		return &model.GameResponse{
			Message: "pending",
		}
	}
}

func (h *GameHandler) StartListening() {

	log.Println("はじまるお")

	for {
		var player1 player
		var player2 player

		player1.conn = <-h.ReadyCh
		player2.conn = <-h.ReadyCh

		log.Println("connきたお")

		player1.UserId = <-h.UserId
		player2.UserId = <-h.UserId

		log.Println("useridきたお")

		player1.PutCard = <-h.PutCard
		player2.PutCard = <-h.PutCard

		log.Println("putCardきたお")

		player1.RoomId = <-h.RoomId
		player2.RoomId = <-h.RoomId

		log.Println("ふたりそろったよ")

		res := h.makeRes(&player1, &player2)

		err := player1.conn.WriteJSON(res)
		if err != nil {
			log.Println("Failed to send message:", err)
		}
		err = player2.conn.WriteJSON(res)
		if err != nil {
			log.Println("Failed to send message:", err)
		}
	}
}
