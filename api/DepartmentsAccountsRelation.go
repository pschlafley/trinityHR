package api

import (
	"net/http"

	"github.com/pschlafley/trinityHR/types"
)

func (s *APIServer) handleCreateDepartmentsAccountsRelation(department_id, account_id int) (*types.DepartmentsAccountsRelation, error) {
	createDeptAccountReq := &types.DepartmentsAccountsRelationReq{DepartmentId: department_id, AccountId: account_id}

	id, err := s.store.CreateDepartmentsAccountsRelation(createDeptAccountReq)

	if err != nil {
		return &types.DepartmentsAccountsRelation{}, err
	}

	DepartmentAccountRelation := types.DepartmentsAccountsRelation{}

	newRelation := DepartmentAccountRelation.NewDepartmentsAccountsRelation(id, createDeptAccountReq.DepartmentId, createDeptAccountReq.AccountId)

	return newRelation, nil
}

func (s *APIServer) handleGetDepartmentsAccountsRelation(w http.ResponseWriter, r *http.Request) error {
	data, err := s.store.GetDepartmentsAccountsRelation()

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, data)
}
