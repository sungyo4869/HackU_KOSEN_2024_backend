package main

import (
	"net/http"

	"github.com/sugyo4869/HackU_KOSEN_2024/handler/router"
)

func main() {
	mux := router.NewRouter()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	srv.ListenAndServe()
}
