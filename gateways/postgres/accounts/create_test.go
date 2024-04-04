package accounts

import (
	"testing"

	"github.com/duarte2025/farmersBank/domain/entities"
	"github.com/jackc/pgx/v5/pgxpool"
)

func TestCreate(t *testing.T) {
	// t.Parallel()

	tests := []struct {
		name    string
		args    entities.Account
		seed    func(t *testing.T, pool *pgxpool.Pool)
		wantErr error
	}{
		{
			name: "create account",
			args: entities.NewAccount("test"),
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			// t.Parallel()
			pgPool := dbPGXPool

			r := NewRepository(pgPool)
			if tt.seed != nil {
				tt.seed(t, pgPool)
			}

			err := r.Create(testCtx, tt.args)
			if err != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}

			// TODO after create FindOne
			// gotFromDB, err := r.FindOne(testCtx, tt.args.ID())
			// require.NoError(t, err)
			// assert.Equal(t, tt.want, gotFromDB)

			pgPool.Close()

		})
	}
}
