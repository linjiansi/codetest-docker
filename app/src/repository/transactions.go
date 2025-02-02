package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/linjiansi/codetest-docker/src/domain/model"
	"github.com/linjiansi/codetest-docker/src/domain/repository"
	"github.com/linjiansi/codetest-docker/src/util"
)

type transactionsRepository struct {
	db *sqlx.DB
}

func NewTransactionsRepository(db *sqlx.DB) repository.TransactionsRepository {
	return &transactionsRepository{db}
}

func (t *transactionsRepository) FetchTotalProductAmount(ctx context.Context, userId int) (int, error) {
	var totalAmount int
	err := t.db.Get(&totalAmount, "SELECT COALESCE(SUM(amount), 0) FROM transactions WHERE user_id = ?", userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil
		} else {
			return 0, util.NewInternalServerError(err)
		}
	}
	return totalAmount, nil
}

func (t *transactionsRepository) InsertTransaction(ctx context.Context, product *model.Product) error {
	tx, err := t.db.BeginTxx(ctx, nil)
	if err != nil {
		return util.NewInternalServerError(err)
	}

	defer func() {
		err = errors.Join(err, tx.Rollback())
	}()

	_, err = tx.ExecContext(ctx, "INSERT INTO transactions (user_id, amount, description) VALUES (?, ?, ?)", product.UserId(), product.Amount(), product.Description())
	if err != nil {
		return util.NewInternalServerError(err)
	}

	if err = tx.Commit(); err != nil {
		return util.NewInternalServerError(err)
	}
	return nil
}
