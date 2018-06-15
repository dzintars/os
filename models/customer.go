package models

// Customer struct
type Customer struct {
	ID        int    `json:"ID"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
}

// ListCustomers function return list of all customers
func ListCustomers() ([]Customer, error) {

	getCustomers := `SELECT
			id,
			name,
			created_at
		FROM os_customers`

	db := dbLoc()
	rows, err := db.Query(getCustomers)
	// ToDo: Handle error
	checkErr(err)

	c := Customer{}
	cs := []Customer{}

	for rows.Next() {
		var (
			id        int
			name      string
			createdAt string
		)

		err := rows.Scan(&id, &name, &createdAt)
		// ToDo: Handle error
		checkErr(err)

		c.ID = id
		c.Name = name
		c.CreatedAt = createdAt

		cs = append(cs, c)
	}
	defer db.Close()

	return cs, nil
}

// GetCustomer function return list of all customers
func GetCustomer(id string) ([]Customer, error) {

	getCustomers := `SELECT
			id,
			name,
			created_at
		FROM os_customers
		WHERE os_customer.id = ?`

	db := dbLoc()
	rows, err := db.Query(getCustomers, id)
	// ToDo: Handle error
	checkErr(err)

	c := Customer{}
	cs := []Customer{}

	for rows.Next() {
		var (
			id        int
			name      string
			createdAt string
		)

		err := rows.Scan(&id, &name, &createdAt)
		// ToDo: Handle error
		checkErr(err)

		c.ID = id
		c.Name = name
		c.CreatedAt = createdAt

		cs = append(cs, c)
	}
	defer db.Close()

	return cs, nil
}
