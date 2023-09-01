package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAdmin(*Account) error
	CreateEmployee(*Account) error
	DeleteEmployee(int) error
	UpdateEmployee(*Account) error
	GetEmployeeByID(int) (*Account, error)
	GetAllAccounts() ([]*Account, error)
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
	return s.createAccountsTable()
}

func (s *PostgresStore) createAccountsTable() (string, error) {
	query := `CREATE TABLE IF NOT EXISTS accounts(
		id serial NOT NULL PRIMARY KEY, 
		account_type varchar(15) NOT NULL,
		full_name varchar(50) NOT NULL,
		email varchar(50) UNIQUE NOT NULL, 
		password varchar(255) UNIQUE NOT NULL,
		created_at timestamp
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

func (s *PostgresStore) CreateEmployee(emp *Account) error {
	query := `INSERT INTO accounts (account_type, full_name, email, password, created_at) VALUES ($1, $2, $3, $4, $5)`

	_, err := s.db.Exec(query, emp.AccountType, emp.FullName, emp.Email, emp.Password, emp.CreatedAt)

	if err != nil {
		return fmt.Errorf("add employee error: %v", err)
	}

	return nil
}

func (s *PostgresStore) CreateAdmin(admin *Account) error {
	query := `INSERT INTO accounts (account_type, full_name, email, password, created_at) VALUES ($1, $2, $3, $4, $5)`

	_, err := s.db.Exec(query, admin.AccountType, admin.FullName, admin.Email, admin.Password, admin.CreatedAt)

	if err != nil {
		return fmt.Errorf("add employee error: %v", err)
	}

	return nil

}

func (s *PostgresStore) DeleteEmployee(id int) error {
	query := `DELETE FROM accounts WHERE id=$1`

	_, err := s.db.Query(query, id)

	if err != nil {
		return fmt.Errorf("unable to delete account with id of %d", id)
	}

	return nil
}

func (s *PostgresStore) UpdateEmployee(*Account) error {
	return nil
}

func (s *PostgresStore) GetEmployeeByID(id int) (*Account, error) {
	query := "select * from accounts where id=$1"

	rows, err := s.db.Query(query, id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoAccount(rows)
	}

	return nil, fmt.Errorf("account %d not found", id)

}

func (s *PostgresStore) GetAllAccounts() ([]*Account, error) {
	rows, err := s.db.Query("select * from accounts")

	if err != nil {
		return nil, err
	}
	var employees []*Account

	for rows.Next() {
		employee, err := scanIntoAccount(rows)

		if err != nil {
			return nil, err
		}

		employees = append(employees, employee)
	}

	return employees, nil
}

func scanIntoAccount(rows *sql.Rows) (*Account, error) {
	employee := &Account{}

	err := rows.Scan(
		&employee.ID,
		&employee.AccountType,
		&employee.FullName,
		&employee.Email,
		&employee.Password,
		&employee.CreatedAt)

	return employee, err
}
