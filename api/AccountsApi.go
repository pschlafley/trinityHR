package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pschlafley/trinityHR/types"
	"golang.org/x/crypto/bcrypt"
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(createEmployeeReq.Password), bcrypt.DefaultCost)

	if err != nil {
		return fmt.Errorf("hashing password error: %v", err)
	}

	accountID, err := s.store.CreateAccount(createEmployeeReq)

	if err != nil {
		return err
	}

	var employee *types.Account

	newEmployee := employee.NewAccount(accountID, string(hashedPassword), createEmployeeReq)

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
