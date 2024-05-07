package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type DatabaseAdapter struct {
	DB *sql.DB
}

func NewDatabaseAdapter(dsn string) *DatabaseAdapter {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error connecting to the database: %s", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging the database: %s", err)
	}

	return &DatabaseAdapter{DB: db}
}

type Employee struct {
	ID       int    `json:"EmployeeID"`
	Username string `json:"Username"`
}

// CreateEmployee adds a new employee record to the database
func (adapter *DatabaseAdapter) CreateEmployee(employee Employee) error {
	_, err := adapter.DB.Exec("INSERT INTO HR (EmployeeID, Username) VALUES (?, ?)", employee.ID, employee.Username)
	return err
}

// GetEmployee retrieves an employee by ID
func (adapter *DatabaseAdapter) GetEmployee(id int) (Employee, error) {
	var emp Employee
	err := adapter.DB.QueryRow("SELECT EmployeeID, Username FROM HR WHERE EmployeeID = ?", id).Scan(&emp.ID, &emp.Username)
	return emp, err
}

// UpdateEmployee updates an existing employee's username
func (adapter *DatabaseAdapter) UpdateEmployee(employee Employee) error {
	_, err := adapter.DB.Exec("UPDATE HR SET Username = ? WHERE EmployeeID = ?", employee.Username, employee.ID)
	return err
}

// DeleteEmployee removes an employee record from the database
func (adapter *DatabaseAdapter) DeleteEmployee(id int) error {
	_, err := adapter.DB.Exec("DELETE FROM HR WHERE EmployeeID = ?", id)
	return err
}
