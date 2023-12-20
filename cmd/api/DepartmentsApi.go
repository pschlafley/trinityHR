package api

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pschlafley/trinityHR/types"
)

func (s *APIServer) handleCreateDepartments(c echo.Context) error {
	var request *types.CreateDepartmentRequest
	var dept *types.Departments

	if err := json.NewDecoder(c.Request().Body).Decode(&request); err != nil {
		return err
	}

	id, err := s.store.CreateDepartment(request)
	if err != nil {
		return err
	}

	newDepartment := dept.NewDepartment(request, id)

	return WriteJSON(c.Response().Writer, http.StatusOK, newDepartment)
}

func (s *APIServer) handleGetDepartments(c echo.Context) error {
	departments, err := s.store.GetDepartments()
	if err != nil {
		return err
	}

	return WriteJSON(c.Response().Writer, http.StatusOK, departments)
}
