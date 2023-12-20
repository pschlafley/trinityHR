package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pschlafley/trinityHR/types"
)

func (s *APIServer) handleGetAllAccounts(c echo.Context) error {
	employees, err := s.store.GetAllAccounts()
	if err != nil {
		return err
	}

	return WriteJSON(c.Response().Writer, http.StatusOK, employees)
}

func (s *APIServer) handleGetAccountById(c echo.Context) error {
	id, err := getIdParam(c)

	if err != nil {
		return fmt.Errorf("invalid id given %d", id)
	}

	employee, getEmployeeErr := s.store.GetAccountByID(id)

	if getEmployeeErr != nil {
		return getEmployeeErr
	}

	return WriteJSON(c.Response().Writer, http.StatusOK, employee)
}

func (s *APIServer) handleCreateAccount(c echo.Context) error {
	var createEmployeeReq *types.CreateAccountRequest

	if err := json.NewDecoder(c.Request().Body).Decode(&createEmployeeReq); err != nil {
		return err
	}

	accountID, err := s.store.CreateAccount(createEmployeeReq)
	if err != nil {
		return err
	}

	var employee *types.Account

	newEmployee, err := employee.NewAccount(accountID, createEmployeeReq.Password, createEmployeeReq)
	if err != nil {
		return err
	}

	var relation *types.DepartmentsAccountsRelationReq = &types.DepartmentsAccountsRelationReq{DepartmentId: createEmployeeReq.Department_id, AccountId: accountID}

	_, relationErr := s.store.CreateDepartmentsAccountsRelation(relation)

	if relationErr != nil {
		return relationErr
	}

	return WriteJSON(c.Response().Writer, http.StatusOK, newEmployee)
}

func (s *APIServer) handleDeleteAccount(c echo.Context) error {
	id, convErr := getIdParam(c)

	if convErr != nil {
		return convErr
	}

	err := s.store.DeleteAccount(id)
	if err != nil {
		return err
	}

	return WriteJSON(c.Response().Writer, http.StatusOK, map[string]int{"deleted": id})
}

// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFeHBpcmVzQXQiOjE1MDAwLCJhY2NvdW50SUQiOjJ9.U9H4OLj__OhQh1m13fb_Z8jcIkOYZ-eEDo6FOQQspz4
func (s *APIServer) handleLogin(c echo.Context) error {
	if c.Request().Method != "POST" {
		return fmt.Errorf("method not allowed %s", c.Request().Method)
	}

	var loginReq *types.AccountLoginReq

	err := json.NewDecoder(c.Request().Body).Decode(&loginReq)
	if err != nil {
		return err
	}

	account, err := s.store.GetAccountByEmail(loginReq.Email)

	if !account.ValidatePassword(loginReq.Password) {
		return fmt.Errorf("not authenticated")
	}

	if err != nil {
		return err
	}

	token, err := createJWT(account)
	if err != nil {
		return err
	}

	resp := types.LoginResponse{
		AccountID: account.AccountID,
		Token:     token,
	}

	return WriteJSON(c.Response().Writer, http.StatusOK, resp)
}
