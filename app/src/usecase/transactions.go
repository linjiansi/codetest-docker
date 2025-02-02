package usecase

import (
	"context"
	"errors"

	"github.com/linjiansi/codetest-docker/src/domain/model"
	"github.com/linjiansi/codetest-docker/src/domain/repository"
	"github.com/linjiansi/codetest-docker/src/usecase/dto"
	"github.com/linjiansi/codetest-docker/src/util"
)

const amountLimit = 1000

type TransactionsUsecase interface {
	ExecTransaction(ctx context.Context, input *dto.TransactionsInput) error
}

type transactionsUsecase struct {
	r repository.TransactionsRepository
}

func NewTransactionsUsecase(r repository.TransactionsRepository) TransactionsUsecase {
	return &transactionsUsecase{r}
}

func (t *transactionsUsecase) ExecTransaction(ctx context.Context, input *dto.TransactionsInput) error {
	totalAmount, err := t.r.FetchTotalProductAmount(ctx, input.UserId)
	if err != nil {
		return err
	}

	if totalAmount+input.Amount > amountLimit {
		return util.NewPaymentError(errors.New("amount limit exceeded"))
	}

	product := model.NewProduct(input.UserId, input.Amount, input.Description)

	return t.r.InsertTransaction(ctx, product)
}
