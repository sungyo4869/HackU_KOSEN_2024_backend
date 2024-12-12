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

	mux.HandleFunc("/healthz", middleware.Recovery(handler.NewHealthzHandler()).ServeHTTP)
	mux.HandleFunc("/login", middleware.Recovery(handler.NewLoginHandler(service.NewUserService(DB))).ServeHTTP)
	mux.HandleFunc("/cards", middleware.Recovery(middleware.Auth(handler.NewCardHandler(*service.NewCardService(DB)))).ServeHTTP)
	mux.HandleFunc("/select", middleware.Recovery(middleware.Auth(handler.NewSelectHandler(*service.NewUserSelectService(DB)))).ServeHTTP)
	
	matchingHandler := handler.NewMatchingHandler(*service.NewSelectedCardService(DB), *service.NewRoomService(DB), *service.NewUserService(DB), *service.NewBattleService(DB))
	mux.HandleFunc("/ws/matching", middleware.Recovery(matchingHandler).ServeHTTP)
	go matchingHandler.StartMatching()

	gameHandler := handler.NewGameHandler(*service.NewSelectedCardService(DB), *service.NewRoomService(DB), *service.NewUserService(DB), *service.NewBattleService(DB), *service.NewUserSelectService(DB))
	mux.HandleFunc("/ws/game", middleware.Recovery(gameHandler).ServeHTTP)
	go gameHandler.SendTurnResult()

	shogunHandler := handler.NewShogunHandler(*service.NewBattleService(DB))
	mux.HandleFunc("/ws/shogun", middleware.Recovery(shogunHandler).ServeHTTP)
	go shogunHandler.SendShogun()

	return mux
}

