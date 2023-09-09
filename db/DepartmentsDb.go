package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/pschlafley/trinityHR/DbTypes"
)

func (s *PostgresStore) createDepartmentsTable() error {
	query := `CREATE TABLE IF NOT EXISTS departments(
		department_id serial NOT NULL PRIMARY KEY, 
		department_name varchar(100) NOT NULL,
		created_at timestamp
	)`

	_, err := s.db.Exec(query)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) CreateDepartment(req *DbTypes.CreateDepartmentRequest) (int, error) {
	query := `INSERT INTO departments (department_name, created_at) VALUES ($1, $2) RETURNING department_id`
	var deptID int

	err := s.db.QueryRow(query, req.DepartmentName, time.Now().UTC()).Scan(&deptID)

	if err != nil {
		return 0, fmt.Errorf("error creating a department: %v", err)
	}

	return deptID, nil
}

func (s *PostgresStore) GetDepartments() ([]*DbTypes.Departments, error) {
	query := `SELECT * FROM departments`

	rows, err := s.db.Query(query)

	if err != nil {
		return nil, fmt.Errorf("error querying all departments: %v", err)
	}

	var departments []*DbTypes.Departments

	for rows.Next() {
		department, err := scanIntoDepartments(rows)

		if err != nil {
			return nil, fmt.Errorf("error scanning into departments: %v", err)
		}

		departments = append(departments, department)
	}

	return departments, nil
}

func scanIntoDepartments(rows *sql.Rows) (*DbTypes.Departments, error) {
	var departments DbTypes.Departments

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
