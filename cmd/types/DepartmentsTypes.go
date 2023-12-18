package types

import "time"

type Departments struct {
	DepartmentID   int       `json:"department_id"`
	DepartmentName string    `json:"department_name"`
	CreatedAt      time.Time `json:"created_at"`
}

type CreateDepartmentRequest struct {
	DepartmentName string `json:"department_name"`
}

func (*Departments) NewDepartment(req *CreateDepartmentRequest, id int) *Departments {
	return &Departments{
		DepartmentID:   id,
		DepartmentName: req.DepartmentName,
		CreatedAt:      time.Now().UTC(),
	}
}
