package handler

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

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
	Player  chan *int64
	RoomId  chan *int64
}

func NewMatchingHandler(SCSvc service.SelectedCardService, RmSvc service.RoomService, UsrSvc service.UserService, BtlSvc service.BattleService) *MatchingHandler {
	return &MatchingHandler{
		SCSvc:   &SCSvc,
		RmSvc:   &RmSvc,
		UsrSvc:  &UsrSvc,
		BtlSvc:  &BtlSvc,
		ReadyCh: make(chan *websocket.Conn, 2),
		Player:  make(chan *int64, 2),
		RoomId:  make(chan *int64, 2),
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
		http.Error(w, "Failed to upgrade to WebSocket", http.StatusInternalServerError)
		return
	}

	defer conn.Close()

	req := model.MatchingUserRequest{}
	err = conn.ReadJSON(&req)
	if err != nil {
		log.Println("Failed to receive json:", err)
		return
	}

	h.Player <- &req.UserId
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

func (h *MatchingHandler) createResponse(userId []int64) (*model.MatchingWSResponse, error) {

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

func (h *MatchingHandler) StartMatching() {
	for {
		conn1 := <-h.ReadyCh
		conn2 := <-h.ReadyCh

		player1 := <-h.Player
		player2 := <-h.Player

		res, err := h.createResponse([]int64{*player1, *player2})
		if err != nil {
			log.Println("Failed to make response:", err)
			return
		}

		roomId := <-h.RoomId

		InitRequest, err := h.NewBattleRequest(*player1, res.Players[0], *roomId)
		if err != nil {
			log.Println("Failed to create battle table request: ", err)
			return
		}
		battle1, err := h.BtlSvc.InitializeBattle(InitRequest)
		if err != nil {
			log.Println("Failed to initialize battle table: ", err)
			return
		}

		InitRequest, err = h.NewBattleRequest(*player2, res.Players[1], *roomId)
		if err != nil {
			log.Println("Failed to create battle table request: ", err)
			return
		}

		battle2, err := h.BtlSvc.InitializeBattle(InitRequest)
		if err != nil {
			log.Println("Failed to initialize battle table:", err)
			return
		}
		
		h.replaceRandomAttributes(&res.Players[0], battle1.RandomAttribute)
		h.replaceRandomAttributes(&res.Players[1], battle2.RandomAttribute)

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

func (h *MatchingHandler) NewBattleRequest(playerId int64, player model.Player, roomId int64) (*model.InitializeBattleRequest, error) {
	var redCardId, blueCardId, greenCardId, kameKameCardId, nankuruCardId, randomCardId int64
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

	if redCardId == 0 || blueCardId == 0 || greenCardId == 0 || kameKameCardId == 0 || nankuruCardId == 0 || randomCardId == 0 {
		return nil, fmt.Errorf("missing card attributes for player: %v", player.Username)
	}

	battleRequest := &model.InitializeBattleRequest{
		UserId:          playerId,
		RoomId:          roomId,
		RedCardId:       sql.NullInt64{Int64: redCardId, Valid: true},
		BlueCardId:      sql.NullInt64{Int64: blueCardId, Valid: true},
		GreenCardId:     sql.NullInt64{Int64: greenCardId, Valid: true},
		KameKameCardId:  sql.NullInt64{Int64: kameKameCardId, Valid: true},
		NankuruCardId:   sql.NullInt64{Int64: nankuruCardId, Valid: true},
		RandomCardId:    sql.NullInt64{Int64: randomCardId, Valid: true},
		RandomAttribute: h.CreateRandomColor(),
		Result:          "pending",
	}

	return battleRequest, nil
}

func (h *MatchingHandler) CreateRandomColor() string {
	rand.Seed(time.Now().UnixNano())

	switch rand.Intn(3) {
	case 0:
		return "red"
	case 1:
		return "blue"
	default:
		return "green"
	}
}

func(h *MatchingHandler) replaceRandomAttributes(player *model.Player, randomAttribute string) {
	for i := 0; i < len(player.SelectedCards); i++ {
		if player.SelectedCards[i].Attribute == "random" {
			player.SelectedCards[i].Attribute = randomAttribute
		}
	}
}
