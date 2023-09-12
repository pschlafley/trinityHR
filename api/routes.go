package api

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pschlafley/trinityHR/db"
)

type APIServer struct {
	listenAddr string
	store      db.Storage
}

func NewAPIServer(listenAddr string, store db.Storage) *APIServer {
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

	router.HandleFunc("/", makeHTTPHandleFunc(handleHomePage))

	router.HandleFunc("/api/accounts/create", makeHTTPHandleFunc(s.handleCreateAccount))

	router.HandleFunc("/api/accounts/{id}", makeHTTPHandleFunc(s.handleGetAccountById))

	router.HandleFunc("/api/accounts", makeHTTPHandleFunc(s.handleGetAllAccounts))

	router.HandleFunc("/api/accounts/delete/{id}", makeHTTPHandleFunc(s.handleDeleteAccount))

	router.HandleFunc("/api/time-off/create", makeHTTPHandleFunc(s.handleCreateTimeOff))

	router.HandleFunc("/api/time-off", makeHTTPHandleFunc(s.handleGetTimeOffRequests))

	router.HandleFunc("/api/account-time-off-relation/create", makeHTTPHandleFunc(s.handleCreateAccountsTimeOffRelationTable))

	router.HandleFunc("/api/account-time-off-relation", makeHTTPHandleFunc(s.handleGetAccountsTimeOffRelationTable))

	router.HandleFunc("/api/departments/create", makeHTTPHandleFunc(s.handleCreateDepartments))

	router.HandleFunc("/api/departments", makeHTTPHandleFunc(s.handleGetDepartments))

	log.Printf("server running at http://localhost%v\n", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

func handleHomePage(w http.ResponseWriter, r *http.Request) error {
	templ := template.Must(template.ParseFiles("views/index.html"))

	return templ.Execute(w, nil)
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
