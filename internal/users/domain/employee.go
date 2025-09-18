package domain

import (
	"context"
	"fmt"
)

type Employee struct {
	ID       string `json:"-"`
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
	City     string `json:"city"`
}

type EmployeesRepository interface {
	Create(ctx context.Context, employee Employee) error
	Get(ctx context.Context, id string) (Employee, error)
}

var (
	ErrEmployeeAlreadyExists = fmt.Errorf("employee already exists")
	ErrEmployeeNotFound      = fmt.Errorf("employee not found")
)
