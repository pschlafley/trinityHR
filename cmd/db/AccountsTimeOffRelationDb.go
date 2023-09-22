package db

import (
	"database/sql"
	"fmt"

	"github.com/pschlafley/trinityHR/types"
)

func (s *PostgresStore) createAccountsTimeOffRelationTable() error {
	query := `CREATE TABLE IF NOT EXISTS accountsTimeOffRelation(
		accountsTimeOffRelation_id serial NOT NULL PRIMARY KEY,
		account_id int REFERENCES accounts(account_id),
		time_off_id int REFERENCES timeOff(time_off_id)
	)`

	_, err := s.db.Exec(query)

	if err != nil {
		return fmt.Errorf("error creating accountsTimeOff Table: %v", err)
	}

	return nil
}

func (s *PostgresStore) CreateAccountsTimeOffRelationTableRow(req *types.AccountsTimeOffRelationRequest) (int, error) {
	query := `INSERT INTO accountsTimeOffRelation (account_id, time_off_id) VALUES ($1, $2) RETURNING accountsTimeOffRelation_id`

	var id int

	err := s.db.QueryRow(query, req.AccountID, req.TimeOffID).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("error inserting into accounts_time_off_relation_table: %v", err)
	}

	return id, nil
}

func (s *PostgresStore) GetAccountsTimeOffRelations() ([]*types.AccountTimeOffRelationQueryData, error) {
	query := `SELECT full_name, email, type, start_date, end_date FROM accountsTimeOffRelation
				JOIN accounts ON accountsTimeOffRelation.account_id = accounts.account_id
				JOIN timeOff ON accountsTimeOffRelation.time_off_id = timeOff.time_off_id;`

	rows, err := s.db.Query(query)

	if err != nil {
		return nil, fmt.Errorf("error fetching from accounts time off relation table: %v", err)
	}

	var accountsTimeOffRelationArr []*types.AccountTimeOffRelationQueryData

	for rows.Next() {
		req, err := scanIntoAccountsTimeOffRelationTable(rows)

		if err != nil {
			return nil, fmt.Errorf("error scanning from accounts time off relation table: %v", err)
		}

		accountsTimeOffRelationArr = append(accountsTimeOffRelationArr, req)
	}

	return accountsTimeOffRelationArr, nil
}

func scanIntoAccountsTimeOffRelationTable(rows *sql.Rows) (*types.AccountTimeOffRelationQueryData, error) {
	accountTimeOff := types.AccountTimeOffRelationQueryData{}

	err := rows.Scan(
		&accountTimeOff.FullName,
		&accountTimeOff.Email,
		&accountTimeOff.Type,
		&accountTimeOff.StartDate,
		&accountTimeOff.EndDate,
	)

	if err != nil {
		return nil, err
	}

	return &accountTimeOff, nil
}
