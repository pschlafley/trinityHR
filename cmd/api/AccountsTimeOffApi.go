package api

import (
	"encoding/json"
	"net/http"

	"github.com/pschlafley/trinityHR/types"
)

func (s *APIServer) handleCreateAccountsTimeOffRelationTable(w http.ResponseWriter, r *http.Request) error {
	accountTimeOffRelationRequest := &types.AccountsTimeOffRelationRequest{}

	if decodeErr := json.NewDecoder(r.Body).Decode(accountTimeOffRelationRequest); decodeErr != nil {
		return decodeErr
	}

	id, dbErr := s.store.CreateAccountsTimeOffRelationTableRow(accountTimeOffRelationRequest)

	if dbErr != nil {
		return dbErr
	}

	accountTimeOffRelationTable := &types.AccountsTimeOffRelationTable{}

	request := accountTimeOffRelationTable.NewAccountsTimeOffRelationTable(id, accountTimeOffRelationRequest.AccountID, accountTimeOffRelationRequest.TimeOffID)

	return WriteJSON(w, http.StatusOK, request)
}

func (s *APIServer) handleGetAccountsTimeOffRelationTable(w http.ResponseWriter, r *http.Request) error {
	response, dbErr := s.store.GetAccountsTimeOffRelations()

	if dbErr != nil {
		return dbErr
	}

	return WriteJSON(w, http.StatusOK, response)
}
