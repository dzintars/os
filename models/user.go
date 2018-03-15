package models

// User struct
type User struct {
	ID   int
	Name string
}

// ListUsers is function to retrieve a full list of all users
func ListUsers() ([]User, error) {
	db := dbLoc()
	rows, err := db.Query("SELECT id, name FROM sys_users")
	if err != nil {
		panic(err.Error())
	}
	usr := User{}
	res := []User{}

	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err.Error())
		}
		usr.ID = id
		usr.Name = name
		res = append(res, usr)
	}
	defer db.Close()
	return res, nil
}
