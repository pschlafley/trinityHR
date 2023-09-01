package main

import (
	"time"
)

const (
	AdminAccount string = "admin"
	UserAccount  string = "user"
)

type CreateAccountRequest struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Account struct {
	ID          int       `json:"id"`
	AccountType string    `json:"account_type"`
	FullName    string    `json:"full_name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	CreatedAt   time.Time `json:"created_at"`
}

func (employee *Account) NewAccount(reqURI, full_name, email, password string) *Account {
	var account *Account

	if reqURI == "/accounts/admins/create" {
		account = &Account{
			AccountType: AdminAccount,
			FullName:    full_name,
			Email:       email,
			Password:    password,
			CreatedAt:   time.Now().UTC(),
		}
	} else if reqURI == "/accounts/employees/create" {
		account = &Account{
			AccountType: UserAccount,
			FullName:    full_name,
			Email:       email,
			Password:    password,
			CreatedAt:   time.Now().UTC(),
		}
	}

	return account
}
