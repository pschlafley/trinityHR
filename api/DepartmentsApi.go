package api

import (
	"net/http"
	"text/template"

	"github.com/pschlafley/trinityHR/types"
)

func (s *APIServer) handleCreateDepartments(w http.ResponseWriter, r *http.Request) error {
	var request *types.CreateDepartmentRequest = &types.CreateDepartmentRequest{}
	var dept *types.Departments = &types.Departments{}

	if err := r.ParseForm(); err != nil {
		return err
	}

	deptName := r.FormValue("departmentName")

	request.DepartmentName = deptName

	id, err := s.store.CreateDepartment(request)

	if err != nil {
		return err
	}

	newDepartment := dept.NewDepartment(request, id)

	return WriteJSON(w, http.StatusOK, newDepartment)
}

func (s *APIServer) handleGetDepartments(w http.ResponseWriter, r *http.Request) error {
	templ := template.Must(template.ParseFiles("views/fragments/departments.html"))

	departments, err := s.store.GetDepartments()

	if err != nil {
		return err
	}

	return templ.ExecuteTemplate(w, "departmenst-list", departments)
}
