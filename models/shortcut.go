package models

// Shortcut are user defined shortcut to favorite apps.
// Order is contstructed by combining account id and number of shortcut sequence.
type Shortcut struct {
	ID            int    `json:"ID"`
	ApplicationID int    `json:"applicationID"`
	Order         int    `json:"order"`
	CreatedAt     string `json:"createdAt"`
	Name          string `json:"name"`
	RelativeURL   string `json:"relativeURL"`
	Bcolor        string `json:"bcolor"`
}

// ListShortcuts retrieve all current user shortcuts
func ListShortcuts() ([]Shortcut, error) {
	getShortcuts := `SELECT
	os_account_shortcuts.id,
	os_account_shortcuts.application_id,
	os_account_shortcuts.order,
	os_account_shortcuts.created_at,
	sys_applications.short_name AS name,
	sys_applications.relative_url AS relative_url,
	sys_applications.background_color AS bcolor
  FROM os_account_shortcuts
  LEFT JOIN sys_applications ON os_account_shortcuts.application_id = sys_applications.id
  ORDER BY os_account_shortcuts.order;`

	db := dbLoc()
	rows, err := db.Query(getShortcuts)
	checkErr(err)

	s := Shortcut{}
	ss := []Shortcut{}

	for rows.Next() {
		var (
			id            int
			applicationID int
			order         int
			createdAt     string
			name          string
			relativeURL   string
			bcolor        string
		)

		err := rows.Scan(&id, &applicationID, &order, &createdAt, &name, &relativeURL, &bcolor)
		checkErr(err)

		s.ID = id
		s.ApplicationID = applicationID
		s.Order = order
		s.CreatedAt = createdAt
		s.Name = name
		s.RelativeURL = relativeURL
		s.Bcolor = bcolor

		ss = append(ss, s)
	}
	defer db.Close()
	return ss, nil
}
