package db

import (
	"database/sql"
	"fmt"

	"github.com/pschlafley/trinityHR/types"
)

func (s *PostgresStore) createAccountsTimeOffRelationTable() error {
	query := `CREATE TABLE IF NOT EXISTS accountsTimeOffRelation(
		id serial NOT NULL PRIMARY KEY,
		account_id int REFERENCES accounts(id),
		time_off_id int REFERENCES timeOff(id)
	)`

	_, err := s.db.Exec(query)

	if err != nil {
		return fmt.Errorf("error creating accountsTimeOff Table: %v", err)
	}

	return nil
}

func (s *PostgresStore) CreateAccountsTimeOffRelationTableRow(req *types.AccountsTimeOffRelationTable) error {
	query := `INSERT INTO accountsTimeOffRelation (account_id, time_off_id) VALUES ($1, $2)`

	_, err := s.db.Exec(query, req.AccountID, req.TimeOffID)

	if err != nil {
		return fmt.Errorf("error inserting into accounts_time_off_relation_table: %v", err)
	}

	return nil
}

func (s *PostgresStore) GetAccountsTimeOffRelations() ([]*types.AccountTimeOffRelationQueryData, error) {
	query := `SELECT full_name, email, type, start_date, end_date FROM accountsTimeOffRelation
				JOIN accounts ON accountsTimeOffRelation.account_id = accounts.id
				JOIN timeOff ON accountsTimeOffRelation.time_off_id = timeOff.id;`

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

	fmt.Print(accountsTimeOffRelationArr)

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
