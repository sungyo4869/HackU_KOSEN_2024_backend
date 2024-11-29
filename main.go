package main

import (
	"log"
	"net/http"

	"github.com/sugyo4869/HackU_KOSEN_2024/handler/router"
	"github.com/sugyo4869/HackU_KOSEN_2024/db"

)

func main() {
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
