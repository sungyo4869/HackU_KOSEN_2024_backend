package main

import (
	"log"
	"net/http"
	"os"

	"github.com/sugyo4869/HackU_KOSEN_2024/db"
	"github.com/sugyo4869/HackU_KOSEN_2024/handler/router"
)

func main() {

	logFile, err := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	
	db, err := db.NewDB()
	if err != nil {
		log.Fatalln("main: err =", err)
	} else {
		log.Println("errなかったぽ")
		err = db.Ping()
		if err != nil {
			log.Println("pingできてないよ, err = ", err)
		} else {
			log.Println("pingできたっぽ")
		}
	}
	defer db.Close()

	mux := router.NewRouter(db)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	srv.ListenAndServe()
}
