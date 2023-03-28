package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func ConnectDb() *sql.DB {
	db, err := sql.Open("mysql", "root:root@/golang")
	if err != nil {
		log.Fatal("error in connct to db", err.Error())
	}
	return db
}
