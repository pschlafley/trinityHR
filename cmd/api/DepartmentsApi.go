package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pschlafley/trinityHR/types"
)

func (s *APIServer) handleCreateDepartments(w http.ResponseWriter, r *http.Request) error {
	var request *types.CreateDepartmentRequest
	var dept *types.Departments

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return err
	}

	fmt.Print(request)

	id, err := s.store.CreateDepartment(request)
	if err != nil {
		return err
	}

	newDepartment := dept.NewDepartment(request, id)

	return WriteJSON(w, http.StatusOK, newDepartment)
}

func (s *APIServer) handleGetDepartments(w http.ResponseWriter, r *http.Request) error {
	departments, err := s.store.GetDepartments()
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, departments)
}
