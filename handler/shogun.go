package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/sugyo4869/HackU_KOSEN_2024/model"
	"github.com/sugyo4869/HackU_KOSEN_2024/service"
)

type ShogunWSHandler struct {
	svc  *service.BattleService
	conn chan *websocket.Conn
	req  chan *model.ShogunRequest
}

func NewShogunHandler(svc service.BattleService) *ShogunWSHandler {
	return &ShogunWSHandler{
		svc:  &svc,
		conn: make(chan *websocket.Conn, 2),
		req: make(chan *model.ShogunRequest, 2),
	}
}

func (h ShogunWSHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade to WebSocket:", err)
		http.Error(w, "Failed to upgrade to WebSocket", http.StatusInternalServerError)
		return
	}

	defer conn.Close()

	req := model.ShogunRequest{}
	err = conn.ReadJSON(&req)
	if err != nil {
		log.Println("Failed to receive json:", err)
		return
	}

	h.req <- &req

	log.Println("りくえすときたよ")
	log.Println(req.RoomId, req.UserId, req.ShogunId)
	
	h.conn <- conn

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

func (h *ShogunWSHandler) createResponse(reqs []model.ShogunRequest) (*model.ShogunResponse, error) {

	err := h.svc.UpdateShogun(reqs[0].UserId, reqs[0].RoomId, reqs[0].ShogunId)
	if err != nil {
		log.Println("failed to update shogun id:", err)
		return nil, err
	}

	err = h.svc.UpdateShogun(reqs[1].UserId, reqs[1].RoomId, reqs[1].ShogunId)
	if err != nil {
		log.Println("failed to update shogun id:", err)
		return nil, err
	}

	var res model.ShogunResponse

	p := model.Shogun{
		RoomId:   reqs[0].RoomId,
		UserId:   reqs[0].UserId,
		ShogunId: reqs[0].ShogunId,
	}
	res.Players = append(res.Players, p)

	p = model.Shogun{
		RoomId:   reqs[1].RoomId,
		UserId:   reqs[1].UserId,
		ShogunId: reqs[1].ShogunId,
	}
	res.Players = append(res.Players, p)

	return &res, nil
}

func (h *ShogunWSHandler) SendShogun() {
	for {
		conn1 := <-h.conn
		conn2 := <-h.conn

		req1 := <- h.req
		req2 := <- h.req

		res, err := h.createResponse([]model.ShogunRequest{*req1, *req2})
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
