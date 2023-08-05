package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
	router.HandleFunc("/employees", makeHTTPHandleFunc(s.handleEmployee))

	router.HandleFunc("/employees/create", makeHTTPHandleFunc(s.handleCreateEmployee))

	router.HandleFunc("/employees/{id}", makeHTTPHandleFunc(s.handleGetEmployee))

	log.Printf("server running at http://localhost%v\n", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleEmployee(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetEmployee(w, r)
	}

	if r.Method == "POST" {
		return s.handleCreateEmployee(w, r)
	}

	if r.Method == "DELETE" {
		return s.handleDeleteEmployee(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleGetEmployee(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	fmt.Println("ID:", id)
	return WriteJSON(w, http.StatusOK, &Employee{})
}

func (s *APIServer) handleCreateEmployee(w http.ResponseWriter, r *http.Request) error {
	createEmployeeReq := &CreateEmployeeRequest{}

	if err := json.NewDecoder(r.Body).Decode(createEmployeeReq); err != nil {
		return err
	}

	employee := NewEmployee(createEmployeeReq.FullName, createEmployeeReq.Email, createEmployeeReq.Password)

	if err := s.store.CreateEmployee(employee); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, employee)
}

func (s *APIServer) handleDeleteEmployee(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// func that returns Encoded JSON data
func WriteJSON(w http.ResponseWriter, status int, value any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(value)
}

type ApiError struct {
	Error string
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
