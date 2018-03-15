package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DBHost  = "127.0.0.1"
	DBPort  = ":3306"
	DBUser  = "root"
	DBPass  = "root"
	DBDbase = "oswee"
)

func dbLoc() (db *sql.DB) {
	connection := fmt.Sprintf("%s:%s@/%s", DBUser, DBPass, DBDbase)
	db, err := sql.Open("mysql", connection)
	if err != nil {
		log.Println("Couldn't connect to" + DBDbase)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}

	return db
}
