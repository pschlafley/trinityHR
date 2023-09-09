package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/pschlafley/trinityHR/types"
)

func (s *PostgresStore) createTimeOffTable() error {
	query := `CREATE TABLE IF NOT EXISTS timeOff(
		time_off_id serial NOT NULL PRIMARY KEY,
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
func (s *PostgresStore) CreateTimeOffRequest(req *types.TimeOffRequest) (int, error) {
	query := `INSERT INTO timeOff (type, start_date, end_date, created_at) VALUES ($1, $2, $3, $4) RETURNING time_off_id`

	var timeOffID int

	err := s.db.QueryRow(query, req.Type, req.StartDate, req.EndDate, time.Now().UTC()).Scan(&timeOffID)

	if err != nil {
		return 0, fmt.Errorf("error inserting into timeOff table: %v", err)
	}

	return timeOffID, nil
}

func (s *PostgresStore) GetTimeOffRequests() ([]*types.TimeOff, error) {
	query := `SELECT * FROM timeOff`

	rows, err := s.db.Query(query)

	if err != nil {
		return nil, fmt.Errorf("error fetching from timeOff table: %v", err)
	}

	var requests []*types.TimeOff

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
func scanIntoTimeOffTable(rows *sql.Rows) (*types.TimeOff, error) {
	timeOff := types.TimeOff{}

	err := rows.Scan(
		&timeOff.TimeOffID,
		&timeOff.Type,
		&timeOff.StartDate,
		&timeOff.EndDate,
		&timeOff.CreatedAt,
	)

	return &timeOff, err
}
