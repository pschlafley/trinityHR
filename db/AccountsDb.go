package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/pschlafley/trinityHR/types"
)

func (s *PostgresStore) createAccountsTable() error {
	query := `CREATE TABLE IF NOT EXISTS accounts(
		id serial NOT NULL PRIMARY KEY, 
		account_type varchar(50) NOT NULL,
		role varchar(100),
		full_name varchar(50) NOT NULL,
		email varchar(50) UNIQUE NOT NULL, 
		password varchar(100) UNIQUE NOT NULL,
		created_at timestamp,
		department_id int REFERENCES departments(id)
	)`

	_, err := s.db.Exec(query)

	if err != nil {
		return err
	}

	return nil

}
func (s *PostgresStore) CreateEmployee(emp *types.CreateAccountRequest) (int, error) {
	query := `INSERT INTO accounts (account_type, role, full_name, email, password, created_at, department_id) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	var accountId int

	err := s.db.QueryRow(query, emp.AccountType, emp.Role, emp.FullName, emp.Email, emp.Password, time.Now().UTC(), emp.Department_id).Scan(&accountId)

	if err != nil {
		return 0, fmt.Errorf("add employee error: %v", err)
	}

	return accountId, nil
}

func (s *PostgresStore) CreateAdmin(admin *types.Account) error {
	query := `INSERT INTO accounts (account_type, role, full_name, email, password, created_at) VALUES ($1, $2, $3, $4, $5)`

	_, err := s.db.Exec(query, admin.AccountType, admin.Role, admin.FullName, admin.Email, admin.Password, admin.CreatedAt, admin.Department_id)

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

func (s *PostgresStore) GetAllAccounts() ([]*types.Account, error) {
	rows, err := s.db.Query("select * from accounts")

	if err != nil {
		return nil, fmt.Errorf("error fetching getAllAccounts: %d", err)
	}

	var accounts []*types.Account

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

func (s *PostgresStore) GetEmployeeByID(id int) (*types.Account, error) {
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
func scanIntoAccount(rows *sql.Rows) (*types.Account, error) {
	employee := types.Account{}

	err := rows.Scan(
		&employee.AccountID,
		&employee.Role,
		&employee.AccountType,
		&employee.FullName,
		&employee.Email,
		&employee.Password,
		&employee.CreatedAt,
		&employee.Department_id,
	)

	return &employee, err
}
