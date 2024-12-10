package router

import (
	"database/sql"
	"net/http"

	"github.com/sugyo4869/HackU_KOSEN_2024/handler"
	"github.com/sugyo4869/HackU_KOSEN_2024/handler/middleware"
	"github.com/sugyo4869/HackU_KOSEN_2024/service"
)

func NewRouter(DB *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", handler.NewHealthzHandler().ServeHTTP)
	mux.HandleFunc("/login", handler.NewLoginHandler(service.NewUserService(DB)).ServeHTTP)
	mux.HandleFunc("/cards", middleware.Auth(handler.NewCardHandler(*service.NewCardService(DB))).ServeHTTP)
	mux.HandleFunc("/select", middleware.Auth(handler.NewUserSelectHandler(*service.NewUserSelectService(DB))).ServeHTTP)
	
	h1 := handler.NewMatchingHandler(*service.NewSelectedCardService(DB), *service.NewRoomService(DB), *service.NewUserService(DB), *service.NewBattleService(DB))
	mux.HandleFunc("/ws/matching", h1.ServeHTTP)
	go h1.StartListening()

	h2 := handler.NewGameHandler(*service.NewSelectedCardService(DB), *service.NewRoomService(DB), *service.NewUserService(DB))
	mux.HandleFunc("/ws/game", h2.ServeHTTP)
	go h2.StartListening()

	return mux
}

