package db

import (
	"database/sql"
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"

	"github.com/pschlafley/trinityHR/types"
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

func (s *PostgresStore) CreateAccount(emp *types.CreateAccountRequest) (int, error) {
	query := `INSERT INTO accounts (account_type, role, full_name, email, password, created_at, department_id) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING account_id`
	hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(emp.Password), bcrypt.DefaultCost)

	if hashErr != nil {
		return 0, hashErr
	}

	var accountId int

	err := s.db.QueryRow(query, emp.AccountType, emp.Role, emp.FullName, emp.Email, hashedPassword, time.Now().UTC(), emp.Department_id).Scan(&accountId)
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

func (s *PostgresStore) GetAllAccounts() ([]*types.AccountsDepartmentsRelationData, error) {
	query := `SELECT account_id, account_type, role, full_name, email, accounts.department_id, department_name FROM accounts JOIN departments on accounts.department_id = departments.department_id ORDER BY account_id ASC`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error fetching getAllAccounts: %d", err)
	}

	var accounts []*types.AccountsDepartmentsRelationData

	for rows.Next() {
		account, err := scanIntoAccountDepartmentsRelationData(rows)
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

func (s *PostgresStore) GetAccountByID(id int) (*types.AccountsDepartmentsRelationData, error) {
	query := "SELECT account_id, account_type, role, full_name, email, accounts.department_id, department_name FROM accounts JOIN departments on accounts.department_id = departments.department_id WHERE account_id=$1"

	rows, err := s.db.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf("error fetching getAccountById: %d", err)
	}

	for rows.Next() {
		return scanIntoAccountDepartmentsRelationData(rows)
	}

	return nil, fmt.Errorf("account %d not found", id)
}

func (s *PostgresStore) GetAccountByEmail(email string) (*types.Account, error) {
	query := "SELECT account_id, password FROM accounts WHERE email = $1"
	rows, err := s.db.Query(query, email)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return getAccountByEmailScan(rows)
	}

	return nil, nil
}

func (s *PostgresStore) GetAccountByJWT(token *jwt.Token) (*types.Account, error) {
	tokenID := token.Claims.(jwt.MapClaims)
	query := `SELECT account_id, account_type, role, full_name, email, created_at, department_id FROM accounts WHERE account_id = $1`

	rows, err := s.db.Query(query, tokenID["accountID"])
	if err != nil {
		return nil, err
	}

	var account *types.Account

	for rows.Next() {
		a, err := scanIntoJWTAccountQuery(rows)
		if err != nil {
			return nil, err
		}

		account = a
	}

	return account, nil
}

func getAccountByEmailScan(rows *sql.Rows) (*types.Account, error) {
	employee := types.Account{}

	err := rows.Scan(
		&employee.AccountID,
		&employee.Password,
	)

	return &employee, err
}

func scanIntoAccount(rows *sql.Rows) (*types.Account, error) {
	employee := types.Account{}

	err := rows.Scan(
		&employee.AccountID,
		&employee.AccountType,
		&employee.Role,
		&employee.FullName,
		&employee.Email,
		&employee.CreatedAt,
		&employee.DepartmentID,
	)

	return &employee, err
}

func scanIntoJWTAccountQuery(rows *sql.Rows) (*types.Account, error) {
	employee := types.Account{}

	err := rows.Scan(
		&employee.AccountID,
		&employee.AccountType,
		&employee.Role,
		&employee.FullName,
		&employee.Email,
		&employee.CreatedAt,
		&employee.DepartmentID,
	)

	return &employee, err
}

func scanIntoAccountDepartmentsRelationData(rows *sql.Rows) (*types.AccountsDepartmentsRelationData, error) {
	employee := types.AccountsDepartmentsRelationData{}

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
