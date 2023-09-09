package DbTypes

import (
	"time"
)

type CreateAccountRequest struct {
	AccountType   string `json:"account_type"`
	Role          string `json:"role"`
	FullName      string `json:"full_name"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	Department_id int    `json:"department_id"`
}

type Account struct {
	AccountID     int       `json:"account_id"`
	AccountType   string    `json:"account_type"`
	Role          string    `json:"role"`
	FullName      string    `json:"full_name"`
	Email         string    `json:"email"`
	Password      string    `json:"password"`
	CreatedAt     time.Time `json:"created_at"`
	Department_id int       `json:"department_id"`
}

type AccountsDepartmentsRelationData struct {
	AccountID      int    `json:"account_id"`
	AccountType    string `json:"account_type"`
	Role           string `json:"role"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
	DepartmentID   int    `json:"department_id"`
	DepartmentName string `json:"department_name"`
}

func (*Account) NewAccount(id int, password string, req *CreateAccountRequest) *Account {
	return &Account{
		AccountID:     id,
		AccountType:   req.AccountType,
		Role:          req.Role,
		FullName:      req.FullName,
		Email:         req.Email,
		Password:      password,
		CreatedAt:     time.Now().UTC(),
		Department_id: req.Department_id,
	}
}
