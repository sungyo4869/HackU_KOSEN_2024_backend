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
	BtlSvc *service.BattleService
	ReadyCh chan *websocket.Conn
	Player chan *int
	RoomId chan *int
}

func NewMatchingHandler(SCSvc service.SelectedCardService, RmSvc service.RoomService, UsrSvc service.UserService, BtlSvc service.BattleService) *MatchingHandler {
	return &MatchingHandler{
		SCSvc: &SCSvc,
		RmSvc: &RmSvc,
		UsrSvc: &UsrSvc,
		BtlSvc: &BtlSvc,
		ReadyCh: make(chan *websocket.Conn, 2),
		Player: make(chan *int, 2),
		RoomId: make(chan *int, 2),
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // オリジン全部許可
	},
}

func (h *MatchingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("アクセスきてるね")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade to WebSocket:", err)
		http.Error(w, "Failed to upgrade to WebSocket", http.StatusInternalServerError)	
		return
	}

	defer conn.Close()

	msg := model.ReadUserResponse{}
	err = conn.ReadJSON(&msg)
	if err != nil{
		log.Println("Failed to receive json:", err)
		return
	}

	log.Println("うけとったよ")

	h.Player <- &msg.UserId
	h.ReadyCh <- conn

	log.Println("チャネルに送ったよ")

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

	h.RoomId <- &send.RoomId

	return &send, nil
}

func(h *MatchingHandler) StartListening() {
	for {
		conn1 := <-h.ReadyCh
		conn2 := <-h.ReadyCh
		log.Println("connきた")

		player1 := <-h.Player
		player2 := <-h.Player
		log.Println("idきた")

		log.Println("2人そろったよ")

		res, err := h.makeRes([]int{*player1, *player2})
		if err != nil {
			log.Println("Failed to make response:", err)
			return
		}

		// インデックスで指定してるけど、attribute==fireなインデクスの.cardIdみたいに指定しないとだめ

		roomId := <- h.RoomId

		battle := &model.InitializeBattleRequest{
			UserId: *player1,
			RoomId: *roomId,
			FireCardId: res.Players[0].SelectedCards[0].CardId,
			WaterCardId: res.Players[0].SelectedCards[1].CardId,
			GrassCardId: res.Players[0].SelectedCards[2].CardId,
			KameKameCardId: res.Players[0].SelectedCards[3].CardId,
			NankuruCardId: res.Players[0].SelectedCards[4].CardId,
			RandomCardId: 1,
			Result: "pending",
		}
		h.BtlSvc.InitializeBattle(battle)

		battle = &model.InitializeBattleRequest{
			UserId: *player2,
			RoomId: *roomId,
			FireCardId: res.Players[1].SelectedCards[0].CardId,
			WaterCardId: res.Players[1].SelectedCards[1].CardId,
			GrassCardId: res.Players[1].SelectedCards[2].CardId,
			KameKameCardId: res.Players[1].SelectedCards[3].CardId,
			NankuruCardId: res.Players[1].SelectedCards[4].CardId,
			RandomCardId: 1,
			Result: "pending",
		}

		err = h.BtlSvc.InitializeBattle(battle)
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

