package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/pschlafley/trinityHR/types"
)

func (s *PostgresStore) createDepartmentsAccountsRelationTable() error {
	query := `CREATE TABLE IF NOT EXISTS departmentsaccountsrelation(
		id serial NOT NULL PRIMARY KEY,
		department_id int NOT NULL REFERENCES departments(department_id),
		account_id int NOT NULL REFERENCES accounts(account_id),
		created_at timestamp NOT NULL
	)`

	_, err := s.db.Exec(query)

	if err != nil {
		return fmt.Errorf("error creating DepartmentsAccountsRelationTable: %v", err)
	}

	return nil
}

func (s *PostgresStore) CreateDepartmentsAccountsRelation(req *types.DepartmentsAccountsRelationReq) (int, error) {
	query := `INSERT INTO departmentsaccountsrelation (department_id, account_id, created_at) VALUES ($1, $2, $3) RETURNING id`
	var id int

	err := s.db.QueryRow(query, req.DepartmentId, req.AccountId, time.Now().UTC()).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("error created DepartmentsAccountsRelation: %v", err)
	}

	return id, nil
}

func (s *PostgresStore) GetDepartmentsAccountsRelation() ([]*types.DepartmentsAccountsRelationQuery, error) {
	query := `SELECT id, department_name, full_name, email, role, account_type FROM departmentsaccountsrelation JOIN departments ON departmentsaccountsrelation.department_id = departments.department_id JOIN accounts ON departmentsaccountsrelation.account_id = accounts.account_id`

	rows, err := s.db.Query(query)

	if err != nil {
		return nil, fmt.Errorf("error fetching from DepartmentsAccountsReltaion: %v", err)
	}

	var departmentsAccountsArr []*types.DepartmentsAccountsRelationQuery

	for rows.Next() {
		data, err := scanIntoDepartmentsAccountsRelationQuery(rows)

		if err != nil {
			return nil, fmt.Errorf("error scanning rows into DepartmentsAccountsRelationQuery: %v", err)
		}

		departmentsAccountsArr = append(departmentsAccountsArr, data)
	}

	return departmentsAccountsArr, nil
}

func (s *PostgresStore) GetDepartmentsAccountsRelationByDepartment([]*types.DepartmentsAccountsRelationQuery, error) {

}

func scanIntoDepartmentsAccountsRelationQuery(rows *sql.Rows) (*types.DepartmentsAccountsRelationQuery, error) {
	data := types.DepartmentsAccountsRelationQuery{}

	err := rows.Scan(
		&data.DepartmentId,
		&data.DepartmentName,
		&data.FullName,
		&data.Email,
		&data.Role,
		&data.AccountType,
	)

	if err != nil {
		return nil, err
	}

	return &data, nil
}
