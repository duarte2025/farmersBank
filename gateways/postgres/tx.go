package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type txKey struct{}

type TxManager struct {
	db *pgxpool.Pool
}

func NewTransactionManager(db *pgxpool.Pool) *TxManager {
	return &TxManager{db: db}
}

// querier should be used when either a transaction or a common connection could be used.
type Querier interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
}

// TransactionFromContext extracts transaction from context.
func TransactionFromContext(ctx context.Context) pgx.Tx {
	if tx, ok := ctx.Value(txKey{}).(pgx.Tx); ok {
		return tx
	}

	return nil
}

func (tx TxManager) querier(ctx context.Context) Querier {
	if conn := TransactionFromContext(ctx); conn != nil {
		return conn
	}

	return tx.db
}

func (tx TxManager) Exec(ctx context.Context, query string, args ...any) (pgconn.CommandTag, error) {
	return tx.querier(ctx).Exec(ctx, query, args...)
}

func (tx TxManager) Query(ctx context.Context, query string, args ...any) (pgx.Rows, error) {
	return tx.querier(ctx).Query(ctx, query, args...)
}

func (tx TxManager) QueryRow(ctx context.Context, query string, args ...any) pgx.Row {
	return tx.querier(ctx).QueryRow(ctx, query, args...)
}

func (tx TxManager) SendBatch(ctx context.Context, batch *pgx.Batch) pgx.BatchResults {
	return tx.querier(ctx).SendBatch(ctx, batch)
}

func (tx TxManager) Close() {
	tx.db.Close()
}
