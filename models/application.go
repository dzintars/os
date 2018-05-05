package models

import help "github.com/oswee/os/helpers"

// Application struct
type Application struct {
	ID          int    `json:"ID"`
	Title       string `json:"title"`
	Visibility  int    `json:"visibility"`
	Bcolor      string `json:"bcolor"`
	PrettyTitle string `json:"prettyTitle"`
	RelativeURL string `json:"relativeURL"`
}

// ListApplications is function to retrieve a full list of all users
func ListApplications(visibility int) ([]Application, error) {
	getApplications := `SELECT id, title, visibility, background_color, relative_url FROM sys_applications WHERE visibility=?`

	db := dbLoc()
	rows, err := db.Query(getApplications, visibility)
	checkErr(err)

	app := Application{}
	res := []Application{}

	for rows.Next() {
		var id int
		var title string
		var visibility int
		var bcolor string
		var relativeURL string

		err := rows.Scan(&id, &title, &visibility, &bcolor, &relativeURL)
		checkErr(err)

		app.ID = id
		app.Title = title
		app.Visibility = visibility
		app.Bcolor = bcolor
		app.PrettyTitle = help.PrettyLinks(title)
		app.RelativeURL = relativeURL

		res = append(res, app)
	}
	defer db.Close()
	return res, nil
}
