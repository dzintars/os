package models

// User struct
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// ListUsers is function to retrieve a full list of all users
func ListUsers() ([]User, error) {
	getUsers := "SELECT id, name FROM sys_users"

	db := dbLoc()
	rows, err := db.Query(getUsers)
	checkErr(err)

	usr := User{}
	res := []User{}

	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		checkErr(err)
		usr.ID = id
		usr.Name = name
		res = append(res, usr)
	}
	defer db.Close()
	return res, nil
}
