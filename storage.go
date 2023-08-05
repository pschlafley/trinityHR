package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateEmployee(*Employee) error
	DeleteEmployee(int) error
	UpdateEmployee(*Employee) error
	GetEmployeeByID(int) (*Employee, error)
}

type PostgresStore struct {
	db *sql.DB
}

func (s *PostgresStore) DropTable() error {
	query := `DROP TABLE employees`

	_, err := s.db.Exec(query)

	if err != nil {
		return err
	}

	fmt.Print("employees table dropped")

	return nil
}

func (s *PostgresStore) Init() (string, error) {
	return s.createEmployeesTable()
}

func (s *PostgresStore) createEmployeesTable() (string, error) {
	query := `CREATE TABLE IF NOT EXISTS employees(
		id serial NOT NULL PRIMARY KEY, 
		fullName varchar(50) NOT NULL,
		email varchar(50) NOT NULL, 
		password varchar(50) NOT NULL,
		create_at timestamp
	)`

	_, err := s.db.Exec(query)

	if err != nil {
		return "Error connecting to database", err
	}

	return "Connected to database successfully", nil

}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=postgres password=psTrinityHR sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	store := &PostgresStore{
		db: db,
	}

	return store, nil
}

func (s *PostgresStore) CreateEmployee(emp *Employee) error {
	query := `INSERT INTO employees (fullName, email, password, create_at) VALUES ($1, $2, $3, $4)`
	res, err := s.db.Query(query, emp.FullName, emp.Email, emp.Password, emp.CreatedAt)

	if err != nil {
		return err
	}

	fmt.Printf("%+v", res)

	return nil
}

func (s *PostgresStore) DeleteEmployee(id int) error {
	return nil
}
func (s *PostgresStore) UpdateEmployee(*Employee) error {
	return nil
}
func (s *PostgresStore) GetEmployeeByID(id int) (*Employee, error) {
	return nil, nil
}
