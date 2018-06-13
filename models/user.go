package models

// User struct
type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	CreatedAt string `json:"createdAt"`
}

// ListUsers is function to retrieve a full list of all users
func ListUsers() ([]User, error) {
	getUsers := `SELECT
		sys_users.id,
		sys_users.name,
		sys_users.username,
		sys_users.created_at
	FROM sys_users;`

	db := dbLoc()
	rows, err := db.Query(getUsers)
	checkErr(err)

	u := User{}
	us := []User{}

	for rows.Next() {

		var (
			id        int
			name      string
			username  string
			createdAt string
		)

		err := rows.Scan(&id, &name, &username, &createdAt)
		checkErr(err)

		u.ID = id
		u.Name = name
		u.Username = username
		u.CreatedAt = createdAt

		us = append(us, u)
	}
	defer db.Close()
	return us, nil
}
