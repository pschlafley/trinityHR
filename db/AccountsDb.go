package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/pschlafley/trinityHR/DbTypes"
)

func (s *PostgresStore) createAccountsTable() error {
	query := `CREATE TABLE IF NOT EXISTS accounts(
		account_id serial NOT NULL PRIMARY KEY, 
		account_type varchar(50) NOT NULL,
		role varchar(100),
		full_name varchar(50) NOT NULL,
		email varchar(50) UNIQUE NOT NULL, 
		password varchar(100) UNIQUE NOT NULL,
		created_at timestamp,
		department_id int REFERENCES departments(department_id)
	)`

	_, err := s.db.Exec(query)

	if err != nil {
		return err
	}

	return nil

}

func (s *PostgresStore) CreateAccount(emp *DbTypes.CreateAccountRequest) (int, error) {
	query := `INSERT INTO accounts (account_type, role, full_name, email, password, created_at, department_id) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING account_id`
	var accountId int

	err := s.db.QueryRow(query, emp.AccountType, emp.Role, emp.FullName, emp.Email, emp.Password, time.Now().UTC(), emp.Department_id).Scan(&accountId)

	if err != nil {
		return 0, fmt.Errorf("add employee error: %v", err)
	}

	return accountId, nil
}

func (s *PostgresStore) DeleteAccount(id int) error {
	query := `DELETE FROM accounts WHERE account_id=$1`

	_, err := s.db.Query(query, id)

	if err != nil {
		return fmt.Errorf("(%v) unable to delete account with id of %d", err, id)
	}

	return nil
}

func (s *PostgresStore) GetAllAccounts() ([]*DbTypes.AccountsDepartmentsRelationData, error) {
	query := `SELECT account_id, account_type, role, full_name, email, accounts.department_id, department_name FROM accounts JOIN departments on accounts.department_id = departments.department_id ORDER BY account_id ASC`

	rows, err := s.db.Query(query)

	if err != nil {
		return nil, fmt.Errorf("error fetching getAllAccounts: %d", err)
	}

	var accounts []*DbTypes.AccountsDepartmentsRelationData

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

func (s *PostgresStore) GetAccountByID(id int) (*DbTypes.AccountsDepartmentsRelationData, error) {
	query := "SELECT account_id, account_type, role, full_name, email, accounts.department_id, department_name FROM accounts JOIN departments on accounts.department_id = departments.department_id WHERE account_id=$1"

	rows, err := s.db.Query(query, id)

	if err != nil {
		return nil, fmt.Errorf("error fetching getAccountById: %d", err)
	}

	for rows.Next() {
		return scanIntoAccount(rows)
	}

	return nil, fmt.Errorf("account %d not found", id)
}
func scanIntoAccount(rows *sql.Rows) (*DbTypes.AccountsDepartmentsRelationData, error) {
	employee := DbTypes.AccountsDepartmentsRelationData{}

	err := rows.Scan(
		&employee.AccountID,
		&employee.AccountType,
		&employee.Role,
		&employee.FullName,
		&employee.Email,
		&employee.DepartmentID,
		&employee.DepartmentName,
	)

	return &employee, err
}
