package models

type Project struct {
	ID       int
	Title    string
	Quantity int
}

func GetProjects() ([]Project, error) {
	db := db_loc()
	rows, err := db.Query("SELECT id, title, quantity FROM order_projects WHERE stakeholder_id='1453'")
	if err != nil {
		panic(err.Error())
	}
	pro := Project{}
	res := []Project{}

	for rows.Next() {
		var id int
		var title string
		var quantity int
		err := rows.Scan(&id, &title, &quantity)
		if err != nil {
			panic(err.Error())
		}
		pro.ID = id
		pro.Title = title
		pro.Quantity = quantity
		res = append(res, pro)
	}
	defer db.Close()
	return res, nil
}
