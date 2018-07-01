package models

import "log"

type message struct {
	ID        int64
	Email     string
	Message   string
	CreatedAt string
}

// MessageCreate function creates a new message in DB
func MessageCreate(email, message string) (lastID int64, err error) {
	messageNew := `INSERT INTO os_messages(email, message) VALUES(?,?)`
	db := dbLoc()
	sql, err := db.Prepare(messageNew)
	if err != nil {
		log.Println(err)
	}
	res, err := sql.Exec(email, message)
	if err != nil {
		log.Fatal(err)
	}

	lastID, err = res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	return lastID, err
}
