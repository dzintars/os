package models

import (
	"database/sql"
	"fmt"

	// MySQL Driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/oswee/os/client/helpers"
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
	helpers.CheckErr(err)

	if err = db.Ping(); err != nil {
		panic(err)
	}

	return db
}
