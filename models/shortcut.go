package models

// Shortcut are user defined shortcut to favorite apps.
// Order is contstructed by combining account id and number of shortcut sequence.
type Shortcut struct {
	ID            int    `json:"ID"`
	ApplicationID int    `json:"applicationID"`
	Order         int    `json:"order"`
	CreatedAt     string `json:"createdAt"`
	Name          string `json:"name"`
	Permalink     string `json:"permalink"`
	Bcolor        string `json:"bcolor"`
	IconName      string `json:"iconName"`
}

// ListShortcuts retrieve all current user shortcuts
func ListShortcuts() ([]Shortcut, error) {
	getShortcuts := `SELECT
	os_account_shortcuts.id,
	os_account_shortcuts.application_id,
	os_account_shortcuts.order,
	os_account_shortcuts.created_at,
	sys_applications.short_name AS name,
	sys_applications.permalink AS permalink,
	sys_applications.background_color AS bcolor,
	sys_applications.icon_name
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
			permalink     string
			bcolor        string
			iconName      string
		)

		err := rows.Scan(&id, &applicationID, &order, &createdAt, &name, &permalink, &bcolor, &iconName)
		checkErr(err)

		s.ID = id
		s.ApplicationID = applicationID
		s.Order = order
		s.CreatedAt = createdAt
		s.Name = name
		s.Permalink = permalink
		s.Bcolor = bcolor
		s.IconName = iconName

		ss = append(ss, s)
	}
	defer db.Close()
	return ss, nil
}
