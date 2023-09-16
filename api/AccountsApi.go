package api

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/pschlafley/trinityHR/types"
	"golang.org/x/crypto/bcrypt"
)

func (s *APIServer) handleGetAllAccounts(w http.ResponseWriter, r *http.Request) error {
	tmpl := template.Must(template.ParseFiles("views/fragments/accounts.html"))

	employees, err := s.store.GetAllAccounts()

	if err != nil {
		return err
	}

	return tmpl.ExecuteTemplate(w, "accounts-list", employees)
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
	templ := template.Must(template.ParseFiles("views/fragments/createAccount.html"))
	createEmployeeReq := &types.CreateAccountRequest{}

	if err := r.ParseForm(); err != nil {
		return err
	}

	deptID, err := strconv.Atoi(r.FormValue("deptID"))

	if err != nil {
		return err
	}

	createEmployeeReq.AccountType = r.FormValue("accountType")
	createEmployeeReq.Role = r.FormValue("role")
	createEmployeeReq.FullName = r.FormValue("fullName")
	createEmployeeReq.Email = r.FormValue("email")
	createEmployeeReq.Password = r.FormValue("password")
	createEmployeeReq.Department_id = deptID

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(createEmployeeReq.Password), bcrypt.DefaultCost)

	if err != nil {
		return fmt.Errorf("hashing password error: %v", err)
	}

	accountID, err := s.store.CreateAccount(createEmployeeReq)

	if err != nil {
		return WriteJSON(w, http.StatusInternalServerError, err.Error())
	}

	var employee *types.Account

	newEmployee := employee.NewAccount(accountID, string(hashedPassword), createEmployeeReq)

	return templ.ExecuteTemplate(w, "created-account", WriteJSON(w, http.StatusOK, newEmployee))
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
