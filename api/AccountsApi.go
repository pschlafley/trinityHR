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
func (s *APIServer) handleGetEmployeeById(w http.ResponseWriter, r *http.Request) error {
	id, err := getIdParam(r)

	if err != nil {
		return fmt.Errorf("invalid id given %d", id)
	}

	employee, getEmployeeErr := s.store.GetEmployeeByID(id)

	if getEmployeeErr != nil {
		return getEmployeeErr
	}

	return WriteJSON(w, http.StatusOK, employee)
}

// func (s *APIServer) handleCreateAdmin(w http.ResponseWriter, r *http.Request) error {
// 	createAdminReq := &types.CreateAccountRequest{}
// 	var reqURI string = r.RequestURI

// 	if err := json.NewDecoder(r.Body).Decode(createAdminReq); err != nil {
// 		return err
// 	}

// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(createAdminReq.Password), bcrypt.DefaultCost)

// 	if err != nil {
// 		return err
// 	}

// 	var admin *types.Account

// 	newAdmin := admin.NewAccount(reqURI, createAdminReq.FullName, createAdminReq.Email, string(hashedPassword), createAdminReq.Role, createAdminReq.Department_id)

// 	if err := s.store.CreateAdmin(newAdmin); err != nil {
// 		return err
// 	}

// 	return WriteJSON(w, http.StatusOK, newAdmin)
// }

func (s *APIServer) handleCreateEmployee(w http.ResponseWriter, r *http.Request) error {
	var createEmployeeReq *types.CreateAccountRequest
	var reqURI string = r.RequestURI

	if err := json.NewDecoder(r.Body).Decode(&createEmployeeReq); err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(createEmployeeReq.Password), bcrypt.DefaultCost)

	if err != nil {
		return fmt.Errorf("hashing password error: %v", err)
	}

	createEmployeeReq.Password = string(hashedPassword)

	accountID, err := s.store.CreateEmployee(createEmployeeReq)

	if err != nil {
		return err
	}

	var employee *types.Account

	newEmployee := employee.NewAccount(accountID, reqURI, string(hashedPassword), createEmployeeReq)

	return WriteJSON(w, http.StatusOK, newEmployee)
}

func (s *APIServer) handleDeleteEmployee(w http.ResponseWriter, r *http.Request) error {
	id, convErr := getIdParam(r)

	if convErr != nil {
		return convErr
	}

	err := s.store.DeleteEmployee(id)

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, map[string]int{"deleted": id})
}
