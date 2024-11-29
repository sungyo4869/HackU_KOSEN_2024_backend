package router

import (
	"net/http"
	"database/sql"
	"github.com/sugyo4869/HackU_KOSEN_2024/handler"
)

func NewRouter(DB *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", handler.NewHealthzHandler().ServeHTTP)
	return mux
}
