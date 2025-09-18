package adapters_test

import (
	"context"
	"employee/internal/users/adapters"
	"employee/internal/users/domain"
	"employee/tests"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/stretchr/testify/require"
)

func TestEmployeesRepository_Create(t *testing.T) {
	ctx := context.Background()

	pg, err := tests.GetPostgresPool(t, ctx)
	require.NoError(t, err)
	defer pg.Close()

	repo := adapters.NewEmployeesRepository(pg)

	tests := []struct {
		name   string
		input  domain.Employee
		want   domain.Employee
		expErr error
	}{
		{
			name: "",
			input: domain.Employee{
				ID:       "00000000-0000-0000-0000-000000000001",
				FullName: "name1",
				Phone:    "8-888-888-88-88",
				City:     "Astana",
			},
			want: domain.Employee{
				ID:       "00000000-0000-0000-0000-000000000001",
				FullName: "name1",
				Phone:    "8-888-888-88-88",
				City:     "Astana",
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
			want: domain.Employee{
				ID:       "00000000-0000-0000-0000-000000000002",
				FullName: "name2",
				Phone:    "8-888-888-88-88",
				City:     "Astana",
			},
			expErr: domain.ErrEmployeeAlreadyExists,
		},
		{
			name: "empty name",
			input: domain.Employee{
				ID:       "00000000-0000-0000-0000-000000000003",
				FullName: "",
				Phone:    "8-888-888-88-88",
				City:     "Almaty",
			},
			want: domain.Employee{
				ID:       "00000000-0000-0000-0000-000000000003",
				FullName: "",
				Phone:    "8-888-888-88-88",
				City:     "Almaty",
			},
			expErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repo.Create(ctx, tt.input)
			require.NoError(t, err)

			got, err := repo.Get(ctx, tt.input.ID)
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}
