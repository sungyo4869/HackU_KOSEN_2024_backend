package router

import (
	"net/http"

	"github.com/sugyo4869/HackU_KOSEN_2024/handler"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", handler.NewHealthzHandler().ServeHTTP)
	return mux
}
