package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/sugyo4869/HackU_KOSEN_2024/model"
	"github.com/sugyo4869/HackU_KOSEN_2024/service"
)

type MatchingHandler struct {
	SCSvc   *service.SelectedCardService
	RmSvc   *service.RoomService
	UsrSvc  *service.UserService
	BtlSvc  *service.BattleService
	ReadyCh chan *websocket.Conn
	Player  chan *int
	RoomId  chan *int
}

func NewMatchingHandler(SCSvc service.SelectedCardService, RmSvc service.RoomService, UsrSvc service.UserService, BtlSvc service.BattleService) *MatchingHandler {
	return &MatchingHandler{
		SCSvc:   &SCSvc,
		RmSvc:   &RmSvc,
		UsrSvc:  &UsrSvc,
		BtlSvc:  &BtlSvc,
		ReadyCh: make(chan *websocket.Conn, 2),
		Player:  make(chan *int, 2),
		RoomId:  make(chan *int, 2),
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
	if err != nil {
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

func (h *MatchingHandler) createRes(userId []int) (*model.MatchingWSResponse, error) {

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

func (h *MatchingHandler) StartListening() {
	for {
		conn1 := <-h.ReadyCh
		conn2 := <-h.ReadyCh
		log.Println("connきた")

		player1 := <-h.Player
		player2 := <-h.Player
		log.Println("idきた")

		log.Println("2人そろったよ")

		res, err := h.createRes([]int{*player1, *player2})
		if err != nil {
			log.Println("Failed to make response:", err)
			return
		}

		roomId := <-h.RoomId

		battle, err := h.CreateBattleRequest(*player1, res.Players[0], *roomId)
		if err != nil {
			log.Println("Failed to create battle table request: ", err)
			return
		}
		err = h.BtlSvc.InitializeBattle(battle)
		if err != nil {
			log.Println("Failed to initialize battle table: ", err)
			return
		}

		battle, err = h.CreateBattleRequest(*player2, res.Players[1], *roomId)
		if err != nil {
			log.Println("Failed to create battle table request: ", err)
			return
		}

		err = h.BtlSvc.InitializeBattle(battle)
		if err != nil {
			log.Println("Failed to initialize battle table:", err)
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

func (h *MatchingHandler) CreateBattleRequest(playerId int, player model.Player, roomId int) (*model.InitializeBattleRequest, error) {
	var redCardId, blueCardId, greenCardId, kameKameCardId, nankuruCardId, randomCardId int
	for _, card := range player.SelectedCards {
		switch card.Attribute {
		case "red":
			redCardId = card.CardId
		case "blue":
			blueCardId = card.CardId
		case "green":
			greenCardId = card.CardId
		case "kamekame":
			kameKameCardId = card.CardId
		case "nankuru":
			nankuruCardId = card.CardId
		case "random":
			randomCardId = card.CardId
		}
	}

	log.Println(redCardId, blueCardId, greenCardId, kameKameCardId, nankuruCardId, randomCardId)

	if redCardId == 0 || blueCardId == 0 || greenCardId == 0 || kameKameCardId == 0 || nankuruCardId == 0 || randomCardId == 0{
		return nil, fmt.Errorf("missing card attributes for player: %v", player.Username)
	}

	// InitializeBattleRequestを生成
	battleRequest := &model.InitializeBattleRequest{
		UserId:         playerId,
		RoomId:         roomId,
		FireCardId:     redCardId,
		WaterCardId:    blueCardId,
		GrassCardId:    greenCardId,
		KameKameCardId: kameKameCardId,
		NankuruCardId:  nankuruCardId,
		RandomCardId:   randomCardId,
		Result:         "pending",
	}

	return battleRequest, nil
}
