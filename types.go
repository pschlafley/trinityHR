package main

import (
	"time"
)

const (
	SuperAdminAccount string = "super_admin"
	AdminAccount      string = "admin"
	UserAccount       string = "user"
)

type CreateAccountRequest struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TimeOffRequest struct {
	Type      string `json:"type"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type AccountsTimeOffRelationRequest struct {
	AccountID int `json:"account_id"`
	TimeOffID int `json:"time_off_id"`
}

type Account struct {
	AccountID   int       `json:"account_id"`
	AccountType string    `json:"account_type"`
	FullName    string    `json:"full_name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	CreatedAt   time.Time `json:"created_at"`
}

type TimeOff struct {
	TimeOffID int       `json:"time_off_id"`
	Type      string    `json:"type"`
	StartDate string    `json:"start_date"`
	EndDate   string    `json:"end_date"`
	CreatedAt time.Time `json:"created_at"`
}

type AccountsTimeOffRelationTable struct {
	ID        int `json:"id"`
	AccountID int `json:"account_id"`
	TimeOffID int `json:"time_off_id"`
}

type AccountTimeOffRelationData struct {
	FullName  string `json:"full_name"`
	Type      string `json:"type"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

func (*Account) NewAccount(reqURI, full_name, email, password string) *Account {
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
	} else if reqURI == "/accounts/super-admin/create" {
		account = &Account{
			AccountType: SuperAdminAccount,
			FullName:    full_name,
			Email:       email,
			Password:    password,
			CreatedAt:   time.Now().UTC(),
		}
	}

	return account
}

func (*TimeOff) NewTimeOffRequest(start_date, end_date, time_off_type string) *TimeOff {
	return &TimeOff{
		Type:      time_off_type,
		StartDate: start_date,
		EndDate:   end_date,
		CreatedAt: time.Now().UTC(),
	}
}

func (*AccountsTimeOffRelationTable) NewAccountsTimeOffRelationTable(account_id, time_off_id int) *AccountsTimeOffRelationTable {
	return &AccountsTimeOffRelationTable{
		AccountID: account_id,
		TimeOffID: time_off_id,
	}
}
