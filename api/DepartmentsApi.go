package api

import (
	"encoding/json"
	"net/http"

	"github.com/pschlafley/trinityHR/types"
)

func (s *APIServer) handleCreateDepartments(w http.ResponseWriter, r *http.Request) error {
	var request *types.CreateDepartmentRequest
	var dept *types.Departments

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
