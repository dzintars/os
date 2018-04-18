package models

import "strings"

// Organization struct
type Organization struct {
	ID      int    `json:"ID"`
	LegalID int64  `json:"legalID"`
	Name    string `json:"name"`
	Form    int    `json:"form"`
}

// ListOrganizations is function to retrieve a full list of all Organizations
func ListOrganizations() ([]Organization, error) {
	getOrganizations := "SELECT id, legal_id, name, form FROM os_organizations"

	db := dbLoc()
	rows, err := db.Query(getOrganizations)
	checkErr(err)

	org := Organization{}
	res := []Organization{}

	for rows.Next() {
		var ID int
		var legalID int64
		var name string
		var form int

		err := rows.Scan(&ID, &legalID, &name, &form)
		checkErr(err)

		org.ID = ID
		org.LegalID = legalID
		org.Name = name
		org.Form = form

		res = append(res, org)
	}
	defer db.Close()
	return res, nil
}

// PrettyLinks is function to make string lower case and replace spaces with dash.
func PrettyLinks(s string) string {
	l := strings.ToLower(s)
	return l
}
