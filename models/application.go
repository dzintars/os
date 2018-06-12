package models

import "github.com/oswee/os/helpers"

// Application struct
type Application struct {
	ID               int    `json:"ID"`
	ShortName        string `json:"shortName"`
	FullName         string `json:"fullName"`
	ShortDescription string `json:"shortDescription"`
	Visibility       int    `json:"visibility"`
	Bcolor           string `json:"bcolor"`
	PrettyTitle      string `json:"prettyTitle"`
	RelativeURL      string `json:"relativeURL"`
}

// ListApplications is function to retrieve a full list of all users
func ListApplications(visibility int) ([]Application, error) {
	getApplications := `SELECT id, short_name, full_name, visibility, background_color, relative_url, short_description FROM sys_applications WHERE visibility=?`

	db := dbLoc()
	rows, err := db.Query(getApplications, visibility)
	checkErr(err)

	app := Application{}
	res := []Application{}

	for rows.Next() {
		var id int
		var shortName string
		var fullName string
		var visibility int
		var bcolor string
		var relativeURL string
		var shortDescription string

		err := rows.Scan(&id, &shortName, &fullName, &visibility, &bcolor, &relativeURL, &shortDescription)
		checkErr(err)

		app.ID = id
		app.ShortName = shortName
		app.FullName = fullName
		app.Visibility = visibility
		app.Bcolor = bcolor
		app.PrettyTitle = helpers.PrettyLinks(shortName)
		app.RelativeURL = relativeURL
		app.ShortDescription = shortDescription

		res = append(res, app)
	}
	defer db.Close()
	return res, nil
}

// ListChildApplications is function to retrieve a list of all applications based on requested parent
func ListChildApplications(parentID int) ([]Application, error) {
	getApplications := `SELECT id, short_name, full_name, visibility, background_color, relative_url, short_description FROM sys_applications WHERE parent_id=?`

	db := dbLoc()
	rows, err := db.Query(getApplications, parentID)
	checkErr(err)

	app := Application{}
	res := []Application{}

	for rows.Next() {
		var id int
		var shortName string
		var fullName string
		var visibility int
		var bcolor string
		var relativeURL string
		var shortDescription string

		err := rows.Scan(&id, &shortName, &fullName, &visibility, &bcolor, &relativeURL, &shortDescription)
		checkErr(err)

		app.ID = id
		app.ShortName = shortName
		app.FullName = fullName
		app.Visibility = visibility
		app.Bcolor = bcolor
		app.PrettyTitle = helpers.PrettyLinks(shortName)
		app.RelativeURL = relativeURL
		app.ShortDescription = shortDescription

		res = append(res, app)
	}
	defer db.Close()
	return res, nil
}
