package models

import "github.com/oswee/os/client/helpers"

// Service struct
type Service struct {
	ID          int     `json:"ID"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	UnitPrice   float32 `json:"unitPrice"`
}

// ListServices is function to retrieve a full list of all Services
func ListServices() ([]Service, error) {

	getServices := `SELECT id, title, description, unit_price FROM os_services`

	db := dbLoc()
	rows, err := db.Query(getServices)
	helpers.CheckErr(err)

	srv := Service{}
	res := []Service{}

	for rows.Next() {
		var ID int
		var title string
		var description string
		var unitPrice float32

		err := rows.Scan(&ID, &title, &description, &unitPrice)
		helpers.CheckErr(err)

		srv.ID = ID
		srv.Title = title
		srv.Description = description
		srv.UnitPrice = unitPrice

		res = append(res, srv)
	}
	defer db.Close()
	return res, nil
}
