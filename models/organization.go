package models

import (
	"github.com/oswee/os/client/helpers"
)

// Organization struct
type Organization struct {
	ID                int     `json:"ID"`
	LegalID           int64   `json:"legalID"`
	Name              string  `json:"name"`
	FirstCharacter    string  `json:"firstCharacter"`
	PrettyLink        string  `json:"prettyLink"`
	Form              int     `json:"form"`
	LegalAddress      string  `json:"legalAddress"`
	VerificationLevel int     `json:"verificationLevel"`
	Color             string  `json:"color"`
	VatRegistrationID string  `json:"vatRegistrationID"`
	Lat               float64 `json:"lat"`
	Lng               float64 `json:"lng"`
}

// ListOrganizations is function to retrieve a full list of all Organizations
func ListOrganizations() ([]Organization, error) {
	getOrganizations := `SELECT id, legal_id, name, form, address_legal, verification_level, color, vat_registration_id, lat, lng FROM os_organizations`

	db := dbLoc()
	rows, err := db.Query(getOrganizations)
	helpers.CheckErr(err)

	org := Organization{}
	res := []Organization{}

	for rows.Next() {

		var (
			ID                int
			legalID           int64
			name              string
			form              int
			addressLegal      string
			verificationLevel int
			color             string
			vatRegistrationID string
			lat               float64
			lng               float64
		)

		err := rows.Scan(&ID, &legalID, &name, &form, &addressLegal, &verificationLevel, &color, &vatRegistrationID, &lat, &lng)
		helpers.CheckErr(err)

		org.ID = ID
		org.LegalID = legalID
		org.Name = name
		org.Form = form
		org.LegalAddress = addressLegal
		org.VerificationLevel = verificationLevel
		org.Color = color
		org.VatRegistrationID = vatRegistrationID
		org.FirstCharacter = string([]rune(name)[0])
		org.PrettyLink = helpers.PrettyLinks(name)
		org.Lat = lat
		org.Lng = lng

		res = append(res, org)
	}
	defer db.Close()
	return res, nil
}

// GetOrganization is a function
func GetOrganization(id string) ([]Organization, error) {
	getOrganizations := `SELECT
			id,
			legal_id,
			name,
			form,
			address_legal,
			verification_level,
			color,
			vat_registration_id
		FROM os_organizations
		WHERE id=?`

	db := dbLoc()
	rows, err := db.Query(getOrganizations, id)
	helpers.CheckErr(err)

	org := Organization{}
	res := []Organization{}

	for rows.Next() {

		var (
			ID                int
			legalID           int64
			name              string
			form              int
			addressLegal      string
			verificationLevel int
			color             string
			vatRegistrationID string
		)

		err := rows.Scan(&ID, &legalID, &name, &form, &addressLegal, &verificationLevel, &color, &vatRegistrationID)
		helpers.CheckErr(err)

		org.ID = ID
		org.LegalID = legalID
		org.Name = name
		org.Form = form
		org.LegalAddress = addressLegal
		org.VerificationLevel = verificationLevel
		org.Color = color
		org.VatRegistrationID = vatRegistrationID
		org.FirstCharacter = string([]rune(name)[0])
		org.PrettyLink = helpers.PrettyLinks(name)

		res = append(res, org)
	}
	defer db.Close()
	return res, nil
}
