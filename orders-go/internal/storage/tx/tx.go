package tx

import (
	"context"
	"database/sql"
)

func Tx(ctx context.Context, db *sql.DB, txFunc func(*sql.Tx) error, opts *sql.TxOptions) error {
	tx, err := db.BeginTx(ctx, opts)
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		}
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	err = txFunc(tx)
	return err
}
