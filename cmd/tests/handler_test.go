package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type (
	Account struct {
		AccountID    int    `json:"account_id"`
		AccountType  string `json:"account_type"`
		Role         string `json:"role"`
		FullName     string `json:"full_name"`
		Email        string `json:"email"`
		Password     string `json:"-"`
		DepartmentID int    `json:"department_id"`
	}
	handler struct {
		db map[string]*Account
	}
)

// func (h *handler) createUser(c echo.Context) error {
// 	u := new(Account)
// 	if err := c.Bind(u); err != nil {
// 		return err
// 	}
// 	return c.JSON(http.StatusCreated, u)
// }

func (h *handler) handleGetAccountById(c echo.Context) error {
	id := c.Param("id")

	user := h.db[id]
	if user == nil {
		return echo.NewHTTPError(http.StatusNotFound, "account not found")
	}
	return c.JSON(http.StatusOK, user)
	// employee, getEmployeeErr := s.store.GetAccountByID(id)

	// if getEmployeeErr != nil {
	// 	return getEmployeeErr
	// }
}

var (
	mockDB   = map[string]*Account{"1": {1, "Employee", "Backend Engineer", "Peyton Schlafley", "peyton@test.com", "test_peyton1", 1}}
	userJSON = `{"account_id":1,"account_type":"Employee","role":"Backend Engineer","full_name":"Peyton Schlafley","email":"peyton@test.com","department_id":1}
`
)

func TestGetUser(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("api/accounts/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	h := &handler{mockDB}

	// Assertions
	if assert.NoError(t, h.handleGetAccountById(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}
}
