package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/pschlafley/trinityHR/types"
)

func (s *PostgresStore) createDepartmentsTable() error {
	_, createTypeErr := s.db.Exec(
		`DO $$ BEGIN 
			IF to_regtype('account') IS NULL THEN	
				CREATE TYPE account AS (
					id int, 
					account_type varchar,
					role varchar,
					full_name varchar,
					email varchar, 
					password varchar,
					created_at timestamp,
					department_id int 
				);
			END IF;
		END $$;`)

	if createTypeErr != nil {
		return createTypeErr
	}

	query := `CREATE TABLE IF NOT EXISTS departments(
		department_id serial NOT NULL PRIMARY KEY, 
		department_name varchar(100) NOT NULL,
		account_data account,
		created_at timestamp
	)`

	_, err := s.db.Exec(query)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) CreateDepartment(req *types.CreateDepartmentRequest) (int, error) {
	query := `INSERT INTO departments (department_name, created_at) VALUES ($1, $2) RETURNING department_id`
	var deptID int

	err := s.db.QueryRow(query, req.DepartmentName, time.Now().UTC()).Scan(&deptID)

	if err != nil {
		return 0, fmt.Errorf("error creating a department: %v", err)
	}

	return deptID, nil
}

func (s *PostgresStore) GetDepartments() ([]*types.Departments, error) {
	query := `SELECT * FROM departments`

	rows, err := s.db.Query(query)

	if err != nil {
		return nil, fmt.Errorf("error querying all departments: %v", err)
	}

	var departments []*types.Departments

	for rows.Next() {
		department, err := scanIntoDepartments(rows)

		if err != nil {
			return nil, fmt.Errorf("error scanning into departments: %v", err)
		}

		departments = append(departments, department)
	}

	return departments, nil
}

func scanIntoDepartments(rows *sql.Rows) (*types.Departments, error) {
	var departments types.Departments

	err := rows.Scan(
		&departments.DepartmentID,
		&departments.DepartmentName,
		&departments.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &departments, err
}
