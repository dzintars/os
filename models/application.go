package models

// Application struct
type Application struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Visibility int    `json:"visibility"`
	Bcolor     string `json:"bcolor"`
}

// ListApplications is function to retrieve a full list of all users
func ListApplications(visibility int) ([]Application, error) {
	getApplications := `SELECT id, title, visibility, background_color FROM sys_applications WHERE visibility=?`

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

		err := rows.Scan(&id, &title, &visibility, &bcolor)
		checkErr(err)

		app.ID = id
		app.Title = title
		app.Visibility = visibility
		app.Bcolor = bcolor

		res = append(res, app)
	}
	defer db.Close()
	return res, nil
}
