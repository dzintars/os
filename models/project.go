package models

// Project represents aany kind of project managed in system
type Project struct {
	ID         int
	CustomerID int
	Name       string
	ServiceID  int
	Units      float32
	Amount     float32
	Cost       float32
	Profit     float32
	CreatedAt  string
}

// ListProjects retrieve list of project
func ListProjects() ([]Project, error) {

	getProjects := `SELECT
		id,
		customer_id,
		name,
		service_id,
		units,
		amount,
		cost,
		profit,
		created_at
	FROM os_projects;`

	db := dbLoc()
	rows, err := db.Query(getProjects)
	if err != nil {
		panic(err.Error())
	}

	p := Project{}
	ps := []Project{}

	for rows.Next() {
		var (
			ID         int
			customerID int
			name       string
			serviceID  int
			units      float32
			amount     float32
			cost       float32
			profit     float32
			createdAt  string
		)

		err := rows.Scan(&ID, &customerID, &name, &serviceID, &units, &amount, &cost, &profit, &createdAt)
		if err != nil {
			panic(err.Error())
		}
		p.ID = ID
		p.CustomerID = customerID
		p.Name = name
		p.ServiceID = serviceID
		p.Units = units
		p.Amount = amount
		p.Cost = cost
		p.Profit = profit
		p.CreatedAt = createdAt

		ps = append(ps, p)
	}
	defer db.Close()
	return ps, nil
}

// ListCustomerProjects retrieve list of project
func ListCustomerProjects(id string) ([]Project, error) {

	getProjects := `SELECT
		id,
		customer_id,
		name,
		service_id,
		units,
		amount,
		cost,
		profit,
		created_at
	FROM os_projects
	WHERE os_projects.customer_id = ?;`

	db := dbLoc()
	rows, err := db.Query(getProjects, id)
	if err != nil {
		panic(err.Error())
	}

	p := Project{}
	ps := []Project{}

	for rows.Next() {
		var (
			ID         int
			customerID int
			name       string
			serviceID  int
			units      float32
			amount     float32
			cost       float32
			profit     float32
			createdAt  string
		)

		err := rows.Scan(&ID, &customerID, &name, &serviceID, &units, &amount, &cost, &profit, &createdAt)
		if err != nil {
			panic(err.Error())
		}
		p.ID = ID
		p.CustomerID = customerID
		p.Name = name
		p.ServiceID = serviceID
		p.Units = units
		p.Amount = amount
		p.Cost = cost
		p.Profit = profit
		p.CreatedAt = createdAt

		ps = append(ps, p)
	}
	defer db.Close()
	return ps, nil
}
