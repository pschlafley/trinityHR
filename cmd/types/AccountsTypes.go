package types

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AccountLoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateAccountRequest struct {
	AccountType   string `json:"account_type"`
	Role          string `json:"role"`
	FullName      string `json:"full_name"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	Department_id int    `json:"department_id"`
}

type Account struct {
	AccountID    int       `json:"account_id"`
	AccountType  string    `json:"account_type"`
	Role         string    `json:"role"`
	FullName     string    `json:"full_name"`
	Email        string    `json:"email"`
	Password     string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
	DepartmentID int       `json:"department_id"`
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

func (*Account) NewAccount(id int, password string, req *CreateAccountRequest) (*Account, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	return &Account{
		AccountID:    id,
		AccountType:  req.AccountType,
		Role:         req.Role,
		FullName:     req.FullName,
		Email:        req.Email,
		Password:     string(hashedPassword),
		CreatedAt:    time.Now().UTC(),
		DepartmentID: req.Department_id,
	}, nil
}
