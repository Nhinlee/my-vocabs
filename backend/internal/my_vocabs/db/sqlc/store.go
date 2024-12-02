package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Store interface {
	Querier
}

type SQLStore struct {
	*Queries
	conn *pgx.Conn
}

// Store Constructor
func NewStore(conn *pgx.Conn) Store {
	return &SQLStore{
		conn:    conn,
		Queries: New(conn),
	}
}

// executes queries function in db transaction & support rollback
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	txOps := pgx.TxOptions{
		// Add TX options here
	}
	tx, err := store.conn.BeginTx(ctx, txOps)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx error: %v, rollback err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit(ctx)
}
