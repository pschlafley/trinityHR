package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/pschlafley/trinityHR/types"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func NewAccount(id, department_id int, accountType, role, fullName, email, password string) (*types.Account, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	return &types.Account{
		AccountID:    id,
		AccountType:  accountType,
		Role:         role,
		FullName:     fullName,
		Email:        email,
		Password:     string(hashedPassword),
		CreatedAt:    time.Now().UTC(),
		DepartmentID: department_id,
	}, nil
}

func TestNewAccount(t *testing.T) {
	acc, err := NewAccount(1, 3, "employee", "admin assistant", "casey", "casey@test.com", "TestingPW")

	assert.Nil(t, err)

	fmt.Printf("%+v\n", acc)
}
