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
	GetEmployeeByID(int) (*Account, error)
	GetAllAccounts() ([]*Account, error)
	CreateTimeOffRequest(*TimeOff) error
	GetTimeOffRequests() ([]*TimeOff, error)
}

type PostgresStore struct {
	db *sql.DB
}

func (s *PostgresStore) DropTable() error {
	dropAccountsTimeOffTable := `DROP TABLE IF EXISTS accountsTimeOff`
	dropTimeOffTable := `DROP TABLE IF EXISTS timeOff`
	dropAccountsTable := `DROP TABLE IF EXISTS accounts`

	if _, accountsTimeOffTableErr := s.db.Exec(dropAccountsTimeOffTable); accountsTimeOffTableErr != nil {
		return accountsTimeOffTableErr
	}

	if _, timeOffErr := s.db.Exec(dropTimeOffTable); timeOffErr != nil {
		return timeOffErr
	}

	if _, accountsTableErr := s.db.Exec(dropAccountsTable); accountsTableErr != nil {
		return accountsTableErr
	}

	fmt.Print("dropped all 3 tables\n")

	return nil
}

func (s *PostgresStore) Init() (string, error) {
	if accountsTableError := s.createAccountsTable(); accountsTableError != nil {
		return "", fmt.Errorf("AccountsTableError: %v", accountsTableError)
	}

	if timeOffTableError := s.createTimeOffTable(); timeOffTableError != nil {
		return "", fmt.Errorf("TimeOffTableError: %v", timeOffTableError)
	}

	if accountsTimeOffTableError := s.createAccountsTimeOffTable(); accountsTimeOffTableError != nil {
		return "", fmt.Errorf("AccountsTimeOffTableError: %v", accountsTimeOffTableError)
	}

	return "Connected Successfully", nil
}

func (s *PostgresStore) createAccountsTable() error {
	query := `CREATE TABLE IF NOT EXISTS accounts(
		id serial NOT NULL PRIMARY KEY, 
		account_type varchar(15) NOT NULL,
		full_name varchar(50) NOT NULL,
		email varchar(50) UNIQUE NOT NULL, 
		password varchar(100) UNIQUE NOT NULL,
		created_at timestamp
	)`

	_, err := s.db.Exec(query)

	if err != nil {
		return err
	}

	return nil

}

func (s *PostgresStore) createTimeOffTable() error {
	query := `CREATE TABLE IF NOT EXISTS timeOff(
		id serial NOT NULL PRIMARY KEY,
		type varchar(50) NOT NULL,
		start_date varchar(50) NOT NULL,
		end_date varchar(50) NOT NULL,
		created_at timestamp NOT NULL
	)`

	_, err := s.db.Exec(query)

	if err != nil {
		return fmt.Errorf("error creating time off table: %v", err)
	}

	return nil
}

func (s *PostgresStore) createAccountsTimeOffTable() error {
	query := `CREATE TABLE IF NOT EXISTS accountsTimeOff(
		id serial NOT NULL PRIMARY KEY,
		account_id int REFERENCES accounts(id),
		time_off_id int REFERENCES timeOff(id)
	)`

	_, err := s.db.Exec(query)

	if err != nil {
		return fmt.Errorf("error creating accountsTimeOff Table: %v", err)
	}

	return nil
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

func (s *PostgresStore) CreateTimeOffRequest(req *TimeOff) error {
	query := `INSERT INTO timeOff (type, start_date, end_date, created_at) VALUES ($1, $2, $3, $4)`

	_, err := s.db.Exec(query, req.Type, req.StartDate, req.EndDate, req.CreatedAt)

	if err != nil {
		return fmt.Errorf("error inserting into timeOff table: %v", err)
	}

	return nil
}

func (s *PostgresStore) GetTimeOffRequests() ([]*TimeOff, error) {
	query := `SELECT * FROM timeOff`

	rows, err := s.db.Query(query)

	if err != nil {
		return nil, fmt.Errorf("error fetching from timeOff table: %v", err)
	}

	var requests []*TimeOff

	for rows.Next() {

		req, err := scanIntoTimeOffTable(rows)
		if err != nil {
			return nil, fmt.Errorf("error fetching timeOff Requests: %v", req)
		}

		requests = append(requests, req)
	}

	if requests == nil {
		return nil, fmt.Errorf("%d timeOff requests found", len(requests))
	}

	return requests, nil
}

func (s *PostgresStore) GetEmployeeByID(id int) (*Account, error) {
	query := "select * from accounts where id=$1"

	rows, err := s.db.Query(query, id)

	if err != nil {
		return nil, fmt.Errorf("error fetching getAccountById: %d", err)
	}

	for rows.Next() {
		return scanIntoAccount(rows)
	}

	return nil, fmt.Errorf("account %d not found", id)

}

func (s *PostgresStore) GetAllAccounts() ([]*Account, error) {
	rows, err := s.db.Query("select * from accounts")

	if err != nil {
		return nil, fmt.Errorf("error fetching getAllAccounts: %d", err)
	}

	var accounts []*Account

	for rows.Next() {
		account, err := scanIntoAccount(rows)

		if err != nil {
			return nil, fmt.Errorf("error scanning into accounts: %d", err)
		}

		accounts = append(accounts, account)
	}

	if accounts == nil {
		return nil, fmt.Errorf("%d accounts found", len(accounts))
	}

	return accounts, nil
}

func scanIntoAccount(rows *sql.Rows) (*Account, error) {
	employee := &Account{}

	err := rows.Scan(
		&employee.AccountID,
		&employee.AccountType,
		&employee.FullName,
		&employee.Email,
		&employee.Password,
		&employee.CreatedAt)

	return employee, err
}

func scanIntoTimeOffTable(rows *sql.Rows) (*TimeOff, error) {
	timeOff := &TimeOff{}

	err := rows.Scan(
		&timeOff.TimeOffID,
		&timeOff.Type,
		&timeOff.StartDate,
		&timeOff.EndDate,
		&timeOff.CreatedAt,
	)

	return timeOff, err
}
