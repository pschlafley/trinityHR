package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/pschlafley/trinityHR/types"

	"github.com/gofor-little/env"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*types.CreateAccountRequest) (int, error)
	DeleteAccount(int) error
	GetAccountByID(int) (*types.AccountsDepartmentsRelationData, error)
	GetAllAccounts() ([]*types.AccountsDepartmentsRelationData, error)
	CreateTimeOffRequest(*types.TimeOffRequest) (int, error)
	GetTimeOffRequests() ([]*types.TimeOff, error)
	CreateAccountsTimeOffRelationTableRow(*types.AccountsTimeOffRelationRequest) (int, error)
	GetAccountsTimeOffRelations() ([]*types.AccountTimeOffRelationQueryData, error)
	CreateDepartment(*types.CreateDepartmentRequest) (int, error)
	GetDepartments() ([]*types.Departments, error)
}

type PostgresStore struct {
	db *sql.DB
}

func (s *PostgresStore) DropTable() error {
	dropDepartmentsTable := `DROP TABLE IF EXISTS departments`
	dropAccountsTimeOffTable := `DROP TABLE IF EXISTS accountsTimeOffRelation`
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

	if _, dropDepartmentsTableErr := s.db.Exec(dropDepartmentsTable); dropDepartmentsTableErr != nil {
		return dropDepartmentsTableErr
	}

	fmt.Print("dropped all 4 tables\n")

	return nil
}

func (s *PostgresStore) Init() (string, error) {
	if departmentsTableError := s.createDepartmentsTable(); departmentsTableError != nil {
		return "", fmt.Errorf("DepartmentsTableError: %v", departmentsTableError)
	}

	if accountsTableError := s.createAccountsTable(); accountsTableError != nil {
		return "", fmt.Errorf("AccountsTableError: %v", accountsTableError)
	}

	if timeOffTableError := s.createTimeOffTable(); timeOffTableError != nil {
		return "", fmt.Errorf("TimeOffTableError: %v", timeOffTableError)
	}

	if accountsTimeOffTableError := s.createAccountsTimeOffRelationTable(); accountsTimeOffTableError != nil {
		return "", fmt.Errorf("AccountsTimeOffTableError: %v", accountsTimeOffTableError)
	}

	return "Connected Successfully", nil
}

func NewPostgresStore() (*PostgresStore, error) {
	if err := env.Load(".env"); err != nil {
		log.Fatal(err)
	}

	DBCONN := env.Get("DBCONN", "")

	db, err := sql.Open("postgres", DBCONN)

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
