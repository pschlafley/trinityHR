package types

import "time"

type DepartmentsAccountsRelation struct {
	ID           int       `json:"id"`
	DepartmentId int       `json:"department_id"`
	AccountId    int       `json:"account_id"`
	CreatedAt    time.Time `json:"created_at"`
}

type DepartmentsAccountsRelationReq struct {
	DepartmentId int `json:"department_id"`
	AccountId    int `json:"account_id"`
}

type DepartmentsAccountsRelationQuery struct {
	DepartmentId   int    `json:"department_id"`
	DepartmentName string `json:"department_name"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
	Role           string `json:"role"`
	AccountType    string `json:"account_type"`
}

func (*DepartmentsAccountsRelation) NewDepartmentsAccountsRelation(id, department_id, account_id int) *DepartmentsAccountsRelation {
	return &DepartmentsAccountsRelation{
		ID:           id,
		DepartmentId: department_id,
		AccountId:    account_id,
		CreatedAt:    time.Now().UTC(),
	}
}
