package api

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pschlafley/trinityHR/types"
)

func (s *APIServer) handleCreateAccountsTimeOffRelationTable(c echo.Context) error {
	accountTimeOffRelationRequest := &types.AccountsTimeOffRelationRequest{}

	if decodeErr := json.NewDecoder(c.Request().Body).Decode(accountTimeOffRelationRequest); decodeErr != nil {
		return decodeErr
	}

	id, dbErr := s.store.CreateAccountsTimeOffRelationTableRow(accountTimeOffRelationRequest)

	if dbErr != nil {
		return dbErr
	}

	accountTimeOffRelationTable := &types.AccountsTimeOffRelationTable{}

	request := accountTimeOffRelationTable.NewAccountsTimeOffRelationTable(id, accountTimeOffRelationRequest.AccountID, accountTimeOffRelationRequest.TimeOffID)

	return WriteJSON(c.Response().Writer, http.StatusOK, request)
}

func (s *APIServer) handleGetAccountsTimeOffRelationTable(c echo.Context) error {
	response, dbErr := s.store.GetAccountsTimeOffRelations()

	if dbErr != nil {
		return dbErr
	}

	return WriteJSON(c.Response().Writer, http.StatusOK, response)
}
