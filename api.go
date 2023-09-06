package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type APIServer struct {
	listenAddr string
	store      Storage
}

func NewAPIServer(listenAddr string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() {
	// with Mux Router we cannot specify if we are doind GET, POST, DELETE, or PUT request so we need to handle this ourselves
	router := mux.NewRouter()
	// MUX's HandleFunc takes a Path string and func(http.ResponseWriter, *http.Request) which is of the type HttpHandler from the net/http package
	// our handleAccount func returns an error which means that it is not of the same type of function that mux's HandleFunc requires
	// So we need to convert our handler func to HttpHandler type

	router.HandleFunc("/accounts/employees/create", makeHTTPHandleFunc(s.handleCreateEmployee))

	router.HandleFunc("/accounts/{id}", makeHTTPHandleFunc(s.handleGetEmployeeById))

	router.HandleFunc("/accounts", makeHTTPHandleFunc(s.handleGetAllAccounts))

	router.HandleFunc("/accounts/employees/delete/{id}", makeHTTPHandleFunc(s.handleDeleteEmployee))

	router.HandleFunc("/accounts/admins/create", makeHTTPHandleFunc(s.handleCreateAdmin))

	router.HandleFunc("/accounts/super-admin/create", makeHTTPHandleFunc(s.handleCreateAdmin))

	router.HandleFunc("/time-off/create", makeHTTPHandleFunc(s.handleCreateTimeOff))

	router.HandleFunc("/time-off", makeHTTPHandleFunc(s.handleGetTimeOffRequests))

	router.HandleFunc("/account-time-off-relation/create", makeHTTPHandleFunc(s.handleCreateAccountsTimeOffRelationTable))

	router.HandleFunc("/account-time-off-relation", makeHTTPHandleFunc(s.handleGetAccountsTimeOffRelationTable))

	log.Printf("server running at http://localhost%v\n", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

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

func (s *APIServer) handleCreateAdmin(w http.ResponseWriter, r *http.Request) error {
	createAdminReq := &CreateAccountRequest{}
	var reqURI string = r.RequestURI

	if err := json.NewDecoder(r.Body).Decode(createAdminReq); err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(createAdminReq.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	var admin *Account

	newAdmin := admin.NewAccount(reqURI, createAdminReq.FullName, createAdminReq.Email, string(hashedPassword))

	if err := s.store.CreateAdmin(newAdmin); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, newAdmin)
}

func (s *APIServer) handleCreateEmployee(w http.ResponseWriter, r *http.Request) error {
	createEmployeeReq := &CreateAccountRequest{}
	var reqURI string = r.RequestURI

	if err := json.NewDecoder(r.Body).Decode(createEmployeeReq); err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(createEmployeeReq.Password), bcrypt.DefaultCost)

	if err != nil {
		return fmt.Errorf("hashing password error: %v", err)
	}

	var employee *Account

	newEmployee := employee.NewAccount(reqURI, createEmployeeReq.FullName, createEmployeeReq.Email, string(hashedPassword))

	if err := s.store.CreateEmployee(newEmployee); err != nil {
		return err
	}

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

func (s *APIServer) handleCreateTimeOff(w http.ResponseWriter, r *http.Request) error {
	timeOffRequest := &TimeOffRequest{}

	if err := json.NewDecoder(r.Body).Decode(&timeOffRequest); err != nil {
		return fmt.Errorf("error decoding timeOffRequest body: %v", err)
	}

	var timeOffReq *TimeOff

	newTimeOffRequest := timeOffReq.NewTimeOffRequest(timeOffRequest.StartDate, timeOffRequest.EndDate, timeOffRequest.Type)

	if err := s.store.CreateTimeOffRequest(newTimeOffRequest); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, newTimeOffRequest)
}

func (s *APIServer) handleGetTimeOffRequests(w http.ResponseWriter, r *http.Request) error {
	requests, err := s.store.GetTimeOffRequests()

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, requests)
}

func (s *APIServer) handleCreateAccountsTimeOffRelationTable(w http.ResponseWriter, r *http.Request) error {
	accountTimeOffRelationRequest := &AccountsTimeOffRelationRequest{}

	if decodeErr := json.NewDecoder(r.Body).Decode(accountTimeOffRelationRequest); decodeErr != nil {
		return decodeErr
	}

	accountTimeOffRelationTable := &AccountsTimeOffRelationTable{}

	request := accountTimeOffRelationTable.NewAccountsTimeOffRelationTable(accountTimeOffRelationRequest.AccountID, accountTimeOffRelationRequest.TimeOffID)

	dbErr := s.store.CreateAccountsTimeOffRelationTableRow(request)

	if dbErr != nil {
		return dbErr
	}

	return WriteJSON(w, http.StatusOK, request)
}

func (s *APIServer) handleGetAccountsTimeOffRelationTable(w http.ResponseWriter, r *http.Request) error {
	response, dbErr := s.store.GetAccountsTimeOffRelations()

	if dbErr != nil {
		return dbErr
	}

	return WriteJSON(w, http.StatusOK, response)
}

// func that returns Encoded JSON data
func WriteJSON(w http.ResponseWriter, status int, value any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(value)
}

type ApiError struct {
	Error string `json:"error"`
}

// function signature of the function that we are using for the MakeHTTPHandleFunc
type apiFunc func(http.ResponseWriter, *http.Request) error

// this function decorates our API func into an HTTP.HandlerFunc(ResponseWriter, Request)
func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	// return a func that takes ResponseWriter, and Request that doesn't return anything and then it handles the Error from the API handler function
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			// handle error here
			// encode the Error to JSON data
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func getIdParam(r *http.Request) (int, error) {
	idStr := mux.Vars(r)["id"]

	id, convertionErr := strconv.Atoi(idStr)

	if convertionErr != nil {
		return 0, fmt.Errorf("invalid id given %s", idStr)
	}

	return id, nil
}
