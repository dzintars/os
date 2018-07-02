package models

// Project represents aany kind of project managed in system
type Project struct {
	ID           int
	CustomerID   int
	CustomerName string
	Name         string
	ServiceID    int
	ServiceTitle string
	Units        float32
	Amount       float32
	Cost         float32
	Profit       float32
	CreatedAt    string
}

// ListProjects retrieve list of project
func ListProjects() ([]Project, error) {

	getProjects := `SELECT
		os_projects.id,
		os_projects.customer_id,
		os_customers.name AS customer_name,
		os_projects.name,
		os_projects.service_id,
		os_services.title AS service_title,
		os_projects.units,
		os_projects.amount,
		os_projects.cost,
		os_projects.profit,
		os_projects.created_at
	FROM os_projects
	LEFT JOIN os_services ON os_projects.service_id = os_services.id
	LEFT JOIN os_customers ON os_projects.customer_id = os_customers.id;`

	db := dbLoc()
	rows, err := db.Query(getProjects)
	if err != nil {
		panic(err.Error())
	}

	p := Project{}
	ps := []Project{}

	for rows.Next() {
		var (
			ID           int
			customerID   int
			customerName string
			name         string
			serviceID    int
			serviceTitle string
			units        float32
			amount       float32
			cost         float32
			profit       float32
			createdAt    string
		)

		err := rows.Scan(&ID, &customerID, &customerName, &name, &serviceID, &serviceTitle, &units, &amount, &cost, &profit, &createdAt)
		if err != nil {
			panic(err.Error())
		}
		p.ID = ID
		p.CustomerID = customerID
		p.CustomerName = customerName
		p.Name = name
		p.ServiceID = serviceID
		p.ServiceTitle = serviceTitle
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
func ListCustomerProjects(customerID string) ([]Project, error) {

	getProjects := `SELECT
		os_projects.id,
		os_projects.customer_id,
		os_customers.name AS customer_name,
		os_projects.name,
		os_projects.service_id,
		os_services.title AS service_title,
		os_projects.units,
		os_projects.amount,
		os_projects.cost,
		os_projects.profit,
		os_projects.created_at
	FROM os_projects
	LEFT JOIN os_services ON os_projects.service_id = os_services.id
	LEFT JOIN os_customers ON os_projects.customer_id = os_customers.id
	WHERE os_projects.customer_id = ?;`

	db := dbLoc()
	rows, err := db.Query(getProjects, customerID)
	if err != nil {
		panic(err.Error())
	}

	p := Project{}
	ps := []Project{}

	for rows.Next() {
		var (
			ID           int
			customerID   int
			customerName string
			name         string
			serviceID    int
			serviceTitle string
			units        float32
			amount       float32
			cost         float32
			profit       float32
			createdAt    string
		)

		err := rows.Scan(&ID, &customerID, &customerName, &name, &serviceID, &serviceTitle, &units, &amount, &cost, &profit, &createdAt)
		if err != nil {
			panic(err.Error())
		}
		p.ID = ID
		p.CustomerID = customerID
		p.CustomerName = customerName
		p.Name = name
		p.ServiceID = serviceID
		p.ServiceTitle = serviceTitle
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
