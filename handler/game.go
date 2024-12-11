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
	USSvc     *service.UserSelectService
	ReadyCh   chan *websocket.Conn
	UserId    chan *int64
	RoomId    chan *int64
	Attribute chan *string
	CardId    chan *int64
}

type player struct {
	conn       *websocket.Conn
	UserId     *int64
	RoomId     *int64
	CardId     *int64
	Attribute  *string
	TurnResult string
	Hp         int
}

func NewGameHandler(SCSvc service.SelectedCardService, RmSvc service.RoomService, UsrSvc service.UserService, BtlSvc service.BattleService, USSvc service.UserSelectService) *GameHandler {
	return &GameHandler{
		RmSvc:     &RmSvc,
		UsrSvc:    &UsrSvc,
		BtlSvc:    &BtlSvc,
		USSvc:     &USSvc,
		ReadyCh:   make(chan *websocket.Conn, 2),
		UserId:    make(chan *int64, 2),
		RoomId:    make(chan *int64, 2),
		Attribute: make(chan *string, 2),
		CardId:    make(chan *int64, 2),
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
		h.CardId <- &msg.CardId
		h.ReadyCh <- conn
	}
}

func (h *GameHandler) CreateRes(player1 *player, player2 *player) *model.GameResponse {

	var res model.GameResponse

	winningRelations := map[string]string{
		"red":      "green", // 日の出 > 門松
		"green":    "blue",  // 門松 > 甘酒
		"blue":     "red",   // 甘酒 > 日の出
		"kamekame": "all",   // かめーかめー攻撃 > なんくるないさー以外のすべて
		"nankuru":  "all",   // なんくるないさー > すべての攻撃を防ぐ
	}

	if *player1.Attribute == *player2.Attribute {
		player1.TurnResult = "draw"
		player2.TurnResult = "draw"
		log.Println("drawだよ", *player1.UserId, ":", *player1.Attribute," ",  *player2.UserId, ":", *player2.Attribute)
	} else if winningRelations[*player1.Attribute] == *player2.Attribute || (*player1.Attribute == "kamekame" && *player2.Attribute != "nankuru") {
		player1.TurnResult = "win"
		player2.TurnResult = "lose"
		log.Println(*player1.UserId, "のかちだよ", *player1.UserId, ":", *player1.Attribute," ",  *player2.UserId, ":", *player2.Attribute)
	} else if winningRelations[*player2.Attribute] == *player1.Attribute || (*player2.Attribute == "kamekame" && *player1.Attribute != "nankuru") {
		player1.TurnResult = "lose"
		player2.TurnResult = "win"
		log.Println(*player2.UserId, "のかちだよ", *player1.UserId, ":", *player1.Attribute," ",  *player2.UserId, ":", *player2.Attribute)

	} else {
		player1.TurnResult = "draw"
		player2.TurnResult = "draw"
		log.Println("drawだよ", *player1.UserId, ":", *player1.Attribute," ",  *player2.UserId, ":", *player2.Attribute)
	}
	

	battle1, err := h.BtlSvc.ReadBattle(*player1.UserId, *player1.RoomId)
	if err != nil {
		log.Println("failed to read battle table:", err)
		return nil
	}
	player1.Hp = battle1.Hp
	if player1.TurnResult == "lose" {
		player1.Hp -= 1
	}

	battle2, err := h.BtlSvc.ReadBattle(*player2.UserId, *player2.RoomId)
	if err != nil {
		log.Println("failed to read battle table:", err)
		return nil
	}
	player2.Hp = battle2.Hp
	if player2.TurnResult == "lose" {
		player2.Hp -= 1
	}

	attribute1, err := h.USSvc.ReadAttribute(*player1.CardId)
	if err != nil {
		log.Println("failed to read card attribute:", err)
		return nil
	}

	player1.Attribute = &attribute1

	attribute2, err := h.USSvc.ReadAttribute(*player2.CardId)
	if err != nil {
		log.Println("failed to read card attribute:", err)
		return nil
	}

	player2.Attribute = &attribute2

	battle1, err = h.BtlSvc.UpdateBattle(*player1.UserId, *player1.RoomId, *player1.Attribute, player1.Hp)
	if err != nil {
		log.Println("failed to update battle table", err)
		return nil
	}
	battle2, err = h.BtlSvc.UpdateBattle(*player2.UserId, *player2.RoomId, *player2.Attribute, player2.Hp)
	if err != nil {
		log.Println("failed to update battle table", err)
		return nil
	}

	if player1.Hp == 0 || ((!battle1.RedCardId.Valid && !battle1.BlueCardId.Valid && !battle1.GreenCardId.Valid && !battle1.KameKameCardId.Valid && !battle1.NankuruCardId.Valid && !battle1.RandomCardId.Valid) && battle1.Hp < battle2.Hp) {
		battle1, err = h.BtlSvc.UpdateResult(*player1.UserId, *player1.RoomId, "lose")
		if err != nil {
			log.Println("failed to update battle result", err)
			return nil
		}
		battle2, err = h.BtlSvc.UpdateResult(*player2.UserId, *player2.RoomId, "win")
		if err != nil {
			log.Println("failed to update battle result", err)
			return nil
		}
	} else if player2.Hp == 0 || ((!battle1.RedCardId.Valid && !battle1.BlueCardId.Valid && !battle1.GreenCardId.Valid && !battle1.KameKameCardId.Valid && !battle1.NankuruCardId.Valid && !battle1.RandomCardId.Valid) && battle1.Hp > battle2.Hp) {
		battle1, err = h.BtlSvc.UpdateResult(*player1.UserId, *player1.RoomId, "win")
		if err != nil {
			log.Println("failed to update battle result:", err)
			return nil
		}
		battle2, err = h.BtlSvc.UpdateResult(*player2.UserId, *player2.RoomId, "lose")
		if err != nil {
			log.Println("failed to update battle result:", err)
			return nil
		}
	}

	result := &model.GameResult{
		UserId:          *player1.UserId,
		SelectAttribute: *player1.Attribute,
		TurnResult:      player1.TurnResult,

		Hp:             battle1.Hp,
		RedCardId:      battle1.RedCardId,
		BlueCardId:     battle1.BlueCardId,
		GreenCardId:    battle1.GreenCardId,
		KameKameCardId: battle1.KameKameCardId,
		NankuruCardId:  battle1.NankuruCardId,
		RandomCardId:   battle1.RandomCardId,
	}
	res.Results = append(res.Results, *result)

	result = &model.GameResult{
		UserId:          *player2.UserId,
		SelectAttribute: *player1.Attribute,
		TurnResult:      player2.TurnResult,

		Hp:             battle2.Hp,
		RedCardId:      battle2.RedCardId,
		BlueCardId:     battle2.BlueCardId,
		GreenCardId:    battle2.GreenCardId,
		KameKameCardId: battle2.KameKameCardId,
		NankuruCardId:  battle2.NankuruCardId,
		RandomCardId:   battle2.RandomCardId,
	}
	res.Results = append(res.Results, *result)

	return &res
}

func (h *GameHandler) StartListening() {
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

		player1.CardId = <-h.CardId
		player2.CardId = <-h.CardId

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