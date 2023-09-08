package types

import (
	"time"
)

type CreateAccountRequest struct {
	Role          string `json:"role"`
	AccountType   string `json:"account_type"`
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

func (*Account) NewAccount(id int, reqURI string, req *CreateAccountRequest) *Account {
	var account *Account

	if reqURI == "/accounts/admins/create" {
		account = &Account{
			AccountID:     id,
			AccountType:   AdminAccount,
			Role:          req.Role,
			FullName:      req.FullName,
			Email:         req.Email,
			Password:      req.Password,
			CreatedAt:     time.Now().UTC(),
			Department_id: req.Department_id,
		}
	} else if reqURI == "/accounts/employees/create" {
		account = &Account{
			AccountID:     id,
			AccountType:   UserAccount,
			Role:          req.Role,
			FullName:      req.FullName,
			Email:         req.Email,
			Password:      req.Password,
			CreatedAt:     time.Now().UTC(),
			Department_id: req.Department_id,
		}
	} else if reqURI == "/accounts/super-admin/create" {
		account = &Account{
			AccountID:     id,
			AccountType:   SuperAdminAccount,
			Role:          req.Role,
			FullName:      req.FullName,
			Email:         req.Email,
			Password:      req.Password,
			CreatedAt:     time.Now().UTC(),
			Department_id: req.Department_id,
		}
	}

	return account
}
