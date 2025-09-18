package adapters

import (
	"context"
	"employee/internal/users/domain"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EmployeesRepository struct {
	pg *pgxpool.Pool
}

func NewEmployeesRepository(pg *pgxpool.Pool) *EmployeesRepository {
	return &EmployeesRepository{pg: pg}
}

func (r *EmployeesRepository) Create(ctx context.Context, employee domain.Employee) error {
	_, err := r.pg.Exec(ctx, "insert into employees (id, full_name, phone_number, city) values ($1, $2, $3, $4)",
		employee.ID, employee.FullName, employee.Phone, employee.City)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
			return domain.ErrEmployeeAlreadyExists
		}
		return err
	}
	return nil
}

func (r *EmployeesRepository) Get(ctx context.Context, id string) (domain.Employee, error) {
	var employee domain.Employee
	err := r.pg.QueryRow(ctx, "select id, full_name, phone_number, city from employees where id=$1", id).
		Scan(&employee.ID, &employee.FullName, &employee.Phone, &employee.City)
	return employee, err
}
