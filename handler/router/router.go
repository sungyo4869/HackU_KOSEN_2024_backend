package router

import (
	"database/sql"
	"net/http"

	"github.com/sugyo4869/HackU_KOSEN_2024/handler"
	"github.com/sugyo4869/HackU_KOSEN_2024/service"
)

func NewRouter(DB *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", handler.NewHealthzHandler().ServeHTTP)
	mux.HandleFunc("/login", handler.NewLoginHandler(service.NewUserService(DB)).ServeHTTP)
	mux.HandleFunc("/cards", handler.NewCardHandler(*service.NewCardService(DB)).ServeHTTP)
	mux.HandleFunc("/selectedCards", handler.NewHandHandler(*service.NewHandService(DB)).ServeHTTP)
	return mux
}
