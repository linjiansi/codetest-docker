package repository

import (
	"context"

	"github.com/linjiansi/codetest-docker/src/domain/model"
)

type TransactionsRepository interface {
	FetchTotalProductAmount(ctx context.Context, userId int) (int, error)
	InsertTransaction(ctx context.Context, product *model.Product) error
}
