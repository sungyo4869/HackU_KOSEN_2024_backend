package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/sugyo4869/HackU_KOSEN_2024/model"
	"github.com/sugyo4869/HackU_KOSEN_2024/service"
)

type GameHandler struct {
	SCSvc     *service.SelectedCardService
	RmSvc     *service.RoomService
	UsrSvc    *service.UserService
	BtlSvc    *service.BattleService
	ReadyCh   chan *websocket.Conn
	UserId    chan *int64
	RoomId    chan *int64
	Attribute chan *string
}

type player struct {
	conn      *websocket.Conn
	UserId    *int64
	RoomId    *int64
	Attribute *string
}

func NewGameHandler(SCSvc service.SelectedCardService, RmSvc service.RoomService, UsrSvc service.UserService, BtlSvc service.BattleService) *GameHandler {
	return &GameHandler{
		RmSvc:     &RmSvc,
		UsrSvc:    &UsrSvc,
		BtlSvc:    &BtlSvc,
		ReadyCh:   make(chan *websocket.Conn, 2),
		UserId:    make(chan *int64, 2),
		RoomId:    make(chan *int64, 2),
		Attribute: make(chan *string, 2),
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
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
				log.Println("WebSocket connection closed by client")
				break
			}

			log.Println("Failed to receive json:", err)
			return
		}

		h.UserId <- &msg.UserId
		h.RoomId <- &msg.RoomId
		h.Attribute <- &msg.Attribute
		h.ReadyCh <- conn

		log.Println("おくりだしました")
	}
}

func (h *GameHandler) CreateRes(player1 *player, player2 *player) *model.GameResponse {

	var res model.GameResponse
	var Result1, Result2 string

	winningRelations := map[string]string{
		"red":      "green", // 日の出 > 門松
		"green":    "water", // 門松 > 甘酒
		"water":    "red",   // 甘酒 > 日の出
		"kamekame": "all",   // かめーかめー攻撃 > なんくるないさー以外のすべて
		"nankuru":  "all",   // なんくるないさー > すべての攻撃を防ぐ
	}

	if player1.Attribute == player2.Attribute {
		Result1 = "draw"
		Result2 = "draw"
	} else if winningRelations[*player1.Attribute] == *player2.Attribute ||
		*player1.Attribute == "kamekame" && *player2.Attribute != "nankuru" {
		Result1 = "win"
		Result2 = "lose"
	} else if winningRelations[*player2.Attribute] == *player1.Attribute ||
		*player2.Attribute == "kamekame" && *player1.Attribute != "nankuru" {
		Result1 = "lose"
		Result2 = "win"
	} else {
		Result1 = "draw"
		Result2 = "draw"
	}

	battle, err := h.BtlSvc.ReadBattle(*player1.UserId, *player1.RoomId)
	if err != nil {
		log.Println("failed to read battle table:", err)
		return nil
	}
	hp := battle.Hp

	if Result1 == "lose" {
		hp = hp - 1
	}

	battle, err = h.BtlSvc.UpdateBattle(*player1.UserId, *player1.RoomId, *player1.Attribute, hp)
	if err != nil {
		log.Println("failed to update battles:",err)
		return nil
	}

	result := &model.GameResult{
		UserId:          *player1.UserId,
		Hp:              battle.Hp,
		SelectAttribute: *player1.Attribute,
		Result:          Result1,
		RedCardId:       battle.RedCardId,
		BlueCardId:      battle.BlueCardId,
		GreenCardId:     battle.GreenCardId,
		KameKameCardId:  battle.KameKameCardId,
		NankuruCardId:   battle.NankuruCardId,
		RandomCardId:    battle.RandomCardId,
	}
	res.Results = append(res.Results, *result)

	battle, _ = h.BtlSvc.ReadBattle(*player2.UserId, *player2.RoomId)
	hp = battle.Hp

	if Result2 == "lose" {
		hp = hp - 1
	}
	battle, err = h.BtlSvc.UpdateBattle(*player2.UserId, *player2.RoomId, *player2.Attribute, hp)
	if err != nil {
		log.Println(err)
		return nil
	}

	result = &model.GameResult{
		UserId:          *player2.UserId,
		Hp:              2,
		SelectAttribute: *player1.Attribute,
		Result:          Result2,
		RedCardId:       battle.RedCardId,
		BlueCardId:      battle.BlueCardId,
		GreenCardId:     battle.GreenCardId,
		KameKameCardId:  battle.KameKameCardId,
		NankuruCardId:   battle.NankuruCardId,
		RandomCardId:    battle.RandomCardId,
	}
	res.Results = append(res.Results, *result)

	return &res
}

func (h *GameHandler) StartListening() {

	log.Println("はじまるお")

	for {
		var player1 player
		var player2 player

		player1.conn = <-h.ReadyCh
		player2.conn = <-h.ReadyCh

		player1.UserId = <-h.UserId
		player2.UserId = <-h.UserId

		player1.Attribute = <-h.Attribute
		player2.Attribute = <-h.Attribute

		player1.RoomId = <-h.RoomId
		player2.RoomId = <-h.RoomId

		log.Println("ふたりそろったよ")
		log.Println(*player1.UserId, *player1.RoomId, *player1.Attribute)
		log.Println(*player2.UserId, *player2.RoomId, *player2.Attribute)

		res := h.CreateRes(&player1, &player2)

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
