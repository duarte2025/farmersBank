package accounts

import (
	"context"
	"fmt"

	"github.com/duarte2025/farmersBank/domain/entities"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r Repository) Create(ctx context.Context, a entities.Account) error {
	const (
		query = `INSERT INTO accounts(id, name, balance) VALUES($1, $2, $3)`
	)

	_, err := r.q.Exec(ctx, query, a.ID(), a.Name(), a.Balance())
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == pgerrcode.UniqueViolation {
			return fmt.Errorf("account already exists")
		}

		return err
	}

	return nil
}
