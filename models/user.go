package models

type User struct {
	ID   int
	Name string
}

func GetUsers() ([]User, error) {
	db := db_loc()
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
