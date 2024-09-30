package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all functions to execute db queries and transactions
type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// execTx executes a function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

// TransferTxParams contains the input parameters of the tranfer transaction
type TranfersTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TranfersTxResult struct {
	Transfer    Transfers `json:"transfer"`
	FromAccount Accounts  `json:"from_account"`
	ToAccount   Accounts  `json:"to_account"`
	FromEntry   Entries   `json:"from_entry"`
	ToEntry     Entries   `json:"to_entry"`
}

// TransferTx performs a monet transfer from one account to the other
// It creates a tranfer record, add account entries, and update accounts' balance within a single database transaction
// func (store *Store) TransferTx(ctx context.Context, arg TranfersTxParams) (TranfersTxResult, error) {
// 	var result TranfersTxResult

// 	err := store.execTx(ctx, func(q *Queries) error {
// 		result.Tranfer, err = q.CreateTransfer(ctx, CreateTransferParams{
// 			FromAccountID: arg.FromAccountID,
// 			ToAccountID:   arg.ToAccountID,
// 			Amount:        arg.Amount,
// 		})
// 	})
// 	return result, err
// }
