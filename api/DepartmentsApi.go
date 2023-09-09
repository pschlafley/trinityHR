package api

import (
	"encoding/json"
	"net/http"

	"github.com/pschlafley/trinityHR/DbTypes"
)

func (s *APIServer) handleCreateDepartments(w http.ResponseWriter, r *http.Request) error {
	var request *DbTypes.CreateDepartmentRequest
	var dept *DbTypes.Departments

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return err
	}

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
