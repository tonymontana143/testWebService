package service

import (
	"context"
	"employee/internal/users/domain"

	"github.com/google/uuid"
)

type EmployeesService interface {
	CreateEmployee(ctx context.Context, employee domain.Employee) error
}

type employeesService struct {
	repo domain.EmployeesRepository
}

func NewEmployeesService(repo domain.EmployeesRepository) EmployeesService {
	return &employeesService{repo: repo}
}

func (s *employeesService) CreateEmployee(ctx context.Context, employee domain.Employee) error {
	employeeID := uuid.NewString()

	return s.repo.Create(ctx, domain.Employee{
		ID:       employeeID,
		FullName: employee.FullName,
		Phone:    employee.Phone,
		City:     employee.City,
	})
}
