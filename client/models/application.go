package models

import (
	"github.com/oswee/os/client/helpers"
)

// Application struct
type Application struct {
	ID               int    `json:"ID"`
	Sequence         int    `json:"sequence"`
	ShortName        string `json:"shortName"`
	FullName         string `json:"fullName"`
	ShortDescription string `json:"shortDescription"`
	Visibility       int    `json:"visibility"`
	Bcolor           string `json:"bcolor"`
	PrettyTitle      string `json:"prettyTitle"`
	Permalink        string `json:"permalink"`
	IconName         string `json:"icon"`
}

// ListApplications is function to retrieve a full list of all users
func ListApplications(visibility int) ([]Application, error) {
	getApplications := `SELECT id, sequence, short_name, full_name, visibility, background_color, permalink, short_description, icon_name FROM sys_applications WHERE visibility=? ORDER BY sys_applications.sequence;`

	db := dbLoc()
	rows, err := db.Query(getApplications, visibility)
	helpers.CheckErr(err)

	app := Application{}
	res := []Application{}

	for rows.Next() {
		var (
			id               int
			sequence         int
			shortName        string
			fullName         string
			visibility       int
			bcolor           string
			permalink        string
			shortDescription string
			iconName         string
		)

		err := rows.Scan(&id, &sequence, &shortName, &fullName, &visibility, &bcolor, &permalink, &shortDescription, &iconName)
		helpers.CheckErr(err)

		app.ID = id
		app.Sequence = sequence
		app.ShortName = shortName
		app.FullName = fullName
		app.Visibility = visibility
		app.Bcolor = bcolor
		app.PrettyTitle = helpers.PrettyLinks(shortName)
		app.Permalink = permalink
		app.ShortDescription = shortDescription
		app.IconName = iconName

		res = append(res, app)
	}
	defer db.Close()
	return res, nil
}

// ListChildApplications is function to retrieve a list of all applications based on requested parent
func ListChildApplications(parentID int) ([]Application, error) {
	getApplications := `SELECT id, sequence, short_name, full_name, visibility, background_color, permalink, short_description, icon_name FROM sys_applications WHERE parent_id=? ORDER BY sys_applications.sequence;`

	db := dbLoc()
	rows, err := db.Query(getApplications, parentID)
	helpers.CheckErr(err)

	app := Application{}
	res := []Application{}

	for rows.Next() {
		var (
			id               int
			sequence         int
			shortName        string
			fullName         string
			visibility       int
			bcolor           string
			permalink        string
			shortDescription string
			iconName         string
		)

		err := rows.Scan(&id, &sequence, &shortName, &fullName, &visibility, &bcolor, &permalink, &shortDescription, &iconName)
		helpers.CheckErr(err)

		app.ID = id
		app.Sequence = sequence
		app.ShortName = shortName
		app.FullName = fullName
		app.Visibility = visibility
		app.Bcolor = bcolor
		app.PrettyTitle = helpers.PrettyLinks(shortName)
		app.Permalink = permalink
		app.ShortDescription = shortDescription
		app.IconName = iconName

		res = append(res, app)
	}
	defer db.Close()
	return res, nil
}