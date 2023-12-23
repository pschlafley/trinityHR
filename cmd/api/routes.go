package api

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

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

func (s *APIServer) Run(app *echo.Echo, server *APIServer) {

	// Middleware
	app.Use(middleware.CORS())
	// app.Use(middleware.CSRF())

	app.Use(middleware.Logger())

	// API Routes
	app.POST("/api/accounts/create", s.handleCreateAccount)

	app.GET("/api/accounts/:id", s.withJWTAuth(s.handleGetAccountById))

	app.GET("/api/accounts", s.withJWTAuth(s.handleGetAllAccounts))

	app.PUT("/api/accounts/delete/:id", s.handleDeleteAccount)

	app.POST("/api/time-off/create", s.handleCreateTimeOff)

	app.GET("/api/time-off", s.handleGetTimeOffRequests)

	app.POST("/api/account-time-off-relation/create", s.handleCreateAccountsTimeOffRelationTable)

	app.GET("/api/account-time-off-relation", s.handleGetAccountsTimeOffRelationTable)

	app.POST("/api/departments/create", s.handleCreateDepartments)

	app.GET("/api/departments", s.handleGetDepartments)

	app.GET("/api/departments-accounts-relation", s.handleGetDepartmentsAccountsRelation)

	app.POST("/login", s.handleLogin)

	fmt.Println("server running at http://localhost:3000/")

	app.Start(server.listenAddr)
}

// func that returns Encoded JSON data
// func WriteJSON(w http.ResponseWriter, status int, value any) error {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(status)
// 	return json.NewEncoder(w).Encode(value)
// }

type ApiError struct {
	Error string `json:"error"`
}

// function signature of the function that we are using for the MakeHTTPHandleFunc
// type apiFunc func(http.ResponseWriter, *http.Request) error

// this function decorates our API func into an HTTP.HandlerFunc(ResponseWriter, Request)
// func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
// 	// return a func that takes ResponseWriter, and Request that doesn't return anything and then it handles the Error from the API handler function
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		if err := f(w, r); err != nil {
// 			// handle error here
// 			// encode the Error to JSON data
// 			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
// 		}
// 	}
// }

func getIdParam(c echo.Context) (int, error) {
	idStr := c.Param("id")

	id, convertionErr := strconv.Atoi(idStr)

	fmt.Print(id)

	if convertionErr != nil {
		return 0, fmt.Errorf("invalid id given %s", idStr)
	}

	return id, nil
}
