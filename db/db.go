package db

import (
	"database/sql"
	"log"
    _ "github.com/go-sql-driver/mysql"
)

func NewDB() (*sql.DB, error) {
	// dsn := "root:-j+Qgve(!#9k@tcp(153.121.44.20)/hacku_kosen_2024?parseTime=true"
	dsn := "root:2PL|RQ)hpevE@tcp(127.0.0.1:3306)/hacku_kosen_2024?parseTime=true"
	
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)

		return nil, err
	}

	return db, nil
}
