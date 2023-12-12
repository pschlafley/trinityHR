package types

import "time"

type DepartmentAccountData struct {
	AccountID    int       `json:"account_id"`
	AccountType  string    `json:"account_type"`
	Role         string    `json:"role"`
	FullName     string    `json:"full_name"`
	Email        string    `json:"email"`
	Password     string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
	DepartmentID int       `json:"department_id"`
}

type Departments struct {
	DepartmentID   int                   `json:"department_id"`
	DepartmentName string                `json:"department_name"`
	AccountData    DepartmentAccountData `json:"account_data"`
	CreatedAt      time.Time             `json:"created_at"`
}

type CreateDepartmentRequest struct {
	DepartmentName string `json:"department_name"`
}

type AddAccountDataReq struct {
	AccountData DepartmentAccountData `json:"account_data"`
}

func (*Departments) AddAccountDataToDepartmentTable(req *AddAccountDataReq, id int) *Departments {
	return &Departments{
		AccountData: req.AccountData,
	}
}

func (*Departments) NewDepartment(req *CreateDepartmentRequest, id int) *Departments {
	return &Departments{
		DepartmentID:   id,
		DepartmentName: req.DepartmentName,
		CreatedAt:      time.Now().UTC(),
	}
}
