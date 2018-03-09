package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DBHost_1  = "127.0.0.1"
	DBPort_1  = ":3306"
	DBUser_1  = "root"
	DBPass_1  = "root"
	DBDbase_1 = "oswee"
)

func db_loc() (db *sql.DB) {
	connection := fmt.Sprintf("%s:%s@/%s", DBUser_1, DBPass_1, DBDbase_1)
	db, err := sql.Open("mysql", connection)
	if err != nil {
		log.Println("Couldn't connect to" + DBDbase_1)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Connected to database", DBHost_1, "@", DBPort_1)

	return db
}
