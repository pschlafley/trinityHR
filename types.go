package main

import (
	"time"
)

type CreateEmployeeRequest struct {
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Employee struct {
	ID        int       `json:"id"`
	FullName  string    `json:"fullName"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewEmployee(fullname string, email string, password string) *Employee {
	return &Employee{
		FullName:  fullname,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now().UTC(),
	}
}
