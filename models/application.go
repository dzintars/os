package models

import (
	"database/sql"
	"fmt"

	"github.com/oswee/os/helpers"
)

// Application struct
type Application struct {
	ID               int    `json:"ID"`
	ParentID         int    `json:"parentID"`
	Sequence         int    `json:"sequence"`
	ShortName        string `json:"shortName"`
	FullName         string `json:"fullName"`
	ShortDescription string `json:"shortDescription"`
	Visibility       int    `json:"visibility"`
	Bcolor           string `json:"bcolor"`
	PrettyTitle      string `json:"prettyTitle"`
	Permalink        string `json:"permalink"`
	IconName         string `json:"icon"`
	Notifications    int32  `json:"notifications"`
	IsNew            bool   `json:"isNew"`
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
	getApplications := `SELECT id, sequence, short_name, full_name, visibility, background_color, permalink, short_description, icon_name, notifications, is_new FROM sys_applications WHERE parent_id=? ORDER BY sys_applications.sequence;`

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
			notifications    int32
			isNew            bool
		)

		err := rows.Scan(&id, &sequence, &shortName, &fullName, &visibility, &bcolor, &permalink, &shortDescription, &iconName, &notifications, &isNew)
		if err != nil {
			fmt.Println(err)
		}

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
		app.Notifications = notifications
		app.IsNew = isNew

		res = append(res, app)
	}
	defer db.Close()
	return res, nil
}

// GetApplication function return requested application
func GetApplication(applicationID int) Application {

	var app Application

	getApplication := `SELECT
			id,
			parent_id,
			short_name,
			full_name,
			background_color,
			permalink,
			short_description,
			icon_name,
			notifications,
			is_new
		FROM sys_applications
		WHERE sys_applications.id = ?`

	db := dbLoc()
	row := db.QueryRow(getApplication, applicationID)
	switch err := row.Scan(&app.ID, &app.ParentID, &app.ShortName, &app.FullName, &app.Bcolor, &app.Permalink, &app.ShortDescription, &app.IconName, &app.Notifications, &app.IsNew); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println("Application:", app, "were returned")
	default:
		panic(err)
	}
	defer db.Close()
	return app
}
