package db

import (
	"fmt"
	"time"

	"github.com/pschlafley/trinityHR/types"
)

func (s *PostgresStore) createDepartmentsTable() error {
	query := `CREATE TABLE IF NOT EXISTS departments(
		id serial NOT NULL PRIMARY KEY, 
		name varchar(100) NOT NULL,
		created_at timestamp
	)`

	_, err := s.db.Exec(query)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) CreateDepartment(req *types.CreateDepartmentRequest) (int, error) {
	query := `INSERT INTO departments (name, created_at) VALUES ($1, $2) RETURNING id`
	var deptID int

	err := s.db.QueryRow(query, req.Name, time.Now().UTC()).Scan(&deptID)

	if err != nil {
		return 0, fmt.Errorf("error creating a department: %v", err)
	}

	return deptID, nil
}
