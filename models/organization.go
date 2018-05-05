package models

import help "github.com/oswee/os/helpers"

// Organization struct
type Organization struct {
	ID                int    `json:"ID"`
	LegalID           int64  `json:"legalID"`
	Name              string `json:"name"`
	FirstCharacter    string `json:"firstCharacter"`
	PrettyLink        string `json:"prettyLink"`
	Form              int    `json:"form"`
	LegalAddress      string `json:"legalAddress"`
	VerificationLevel int    `json:"verificationLevel"`
	Color             string `json:"color"`
	VatRegistrationID string `json:"vatRegistrationID"`
}

// ListOrganizations is function to retrieve a full list of all Organizations
func ListOrganizations() ([]Organization, error) {
	getOrganizations := `SELECT id, legal_id, name, form, address_legal, verification_level, color, vat_registration_id FROM os_organizations`

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
		var addressLegal string
		var verificationLevel int
		var color string
		var vatRegistrationID string

		err := rows.Scan(&ID, &legalID, &name, &form, &addressLegal, &verificationLevel, &color, &vatRegistrationID)
		checkErr(err)

		org.ID = ID
		org.LegalID = legalID
		org.Name = name
		org.Form = form
		org.LegalAddress = addressLegal
		org.VerificationLevel = verificationLevel
		org.Color = color
		org.VatRegistrationID = vatRegistrationID
		org.FirstCharacter = string([]rune(name)[0])
		org.PrettyLink = help.PrettyLinks(name)

		res = append(res, org)
	}
	defer db.Close()
	return res, nil
}
