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
	mux.HandleFunc("/selected-cards", middleware.Auth(handler.NewHandHandler(*service.NewUserSelectedService(DB))).ServeHTTP)
	
	h := handler.NewMatchingHandler(service.NewSelectedCardService(DB), service.NewRoomService(DB))
	mux.HandleFunc("/ws/matching", h.ServeHTTP)
	go h.StartListening()

	return mux
}
