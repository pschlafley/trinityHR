package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pschlafley/trinityHR/types"
)

func (s *APIServer) handleGetAllAccounts(w http.ResponseWriter, r *http.Request) error {
	employees, err := s.store.GetAllAccounts()

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, employees)
}

func (s *APIServer) handleGetAccountById(w http.ResponseWriter, r *http.Request) error {
	id, err := getIdParam(r)

	if err != nil {
		return fmt.Errorf("invalid id given %d", id)
	}

	employee, getEmployeeErr := s.store.GetAccountByID(id)

	if getEmployeeErr != nil {
		return getEmployeeErr
	}

	return WriteJSON(w, http.StatusOK, employee)
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	var createEmployeeReq *types.CreateAccountRequest

	if err := json.NewDecoder(r.Body).Decode(&createEmployeeReq); err != nil {
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

	return WriteJSON(w, http.StatusOK, newEmployee)
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	id, convErr := getIdParam(r)

	if convErr != nil {
		return convErr
	}

	err := s.store.DeleteAccount(id)

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, map[string]int{"deleted": id})
}

func (s *APIServer) handleLogin(w http.ResponseWriter, r *http.Request) error {

	var loginReq *types.AccountLoginReq

	err := json.NewDecoder(r.Body).Decode(&loginReq)

	if err != nil {
		return fmt.Errorf("error decoding login: %v", err)
	}

	account, err := s.store.GetAccountByEmail(loginReq.Email)

	if err != nil {
		return err
	}

	token, err := createJWT(account)

	if err != nil {
		return err
	}

	fmt.Print("JWT token: ", token)

	return nil
}
