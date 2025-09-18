package service

import (
	"context"
	"employee/internal/users/adapters"
	"employee/internal/users/domain"
	"employee/tests"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/stretchr/testify/require"
)

func Test_employeesService_CreateEmployee(t *testing.T) {
	ctx := context.Background()

	pg, err := tests.GetPostgresPool(t, ctx)
	require.NoError(t, err)
	defer pg.Close()

	repo := adapters.NewEmployeesRepository(pg)
	svc := NewEmployeesService(repo)
	tests := []struct {
		name   string
		input  domain.Employee
		expErr error
	}{
		{
			name: "",
			input: domain.Employee{
				ID:       "00000000-0000-0000-0000-000000000001",
				FullName: "name1",
				Phone:    "8-888-888-88-88",
				City:     "moscow",
			},
			expErr: nil,
		},
		{
			name: "name2",
			input: domain.Employee{
				ID:       "00000000-0000-0000-0000-000000000002",
				FullName: "name2",
				Phone:    "8-888-888-88-88",
				City:     "Astana",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := svc.CreateEmployee(ctx, tt.input)
			require.NoError(t, err)
		})
	}
}
