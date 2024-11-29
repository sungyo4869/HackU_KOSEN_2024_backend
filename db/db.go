package db

import (
	"database/sql"
	"log"
    _ "github.com/go-sql-driver/mysql"
)

func NewDB() (*sql.DB, error) {
	dsn := "root:pass@tcp(127.0.0.1:3306)/testdb?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)

		return nil, err
	}

	return db, nil
}
