package handler

import (
	"errors"
	"net/http"

	"github.com/go-fuego/fuego"
	"github.com/linjiansi/codetest-docker/src/handler/request"
	"github.com/linjiansi/codetest-docker/src/handler/response"
	"github.com/linjiansi/codetest-docker/src/usecase"
	"github.com/linjiansi/codetest-docker/src/usecase/dto"
	"github.com/linjiansi/codetest-docker/src/util"
)

type TransactionsHandler interface {
	Transactions(c fuego.ContextWithBody[request.Transaction]) (response.Transactions, error)
}

type transactionsHandler struct {
	u usecase.TransactionsUsecase
}

func NewTransactionsHandler(u usecase.TransactionsUsecase) TransactionsHandler {
	return &transactionsHandler{u}
}

func (t *transactionsHandler) Transactions(c fuego.ContextWithBody[request.Transaction]) (response.Transactions, error) {
	userId, err := util.GetUserId(c.Context())
	if err != nil {
		return response.Transactions{}, err
	}
	body, err := c.Body()
	if err != nil {
		return response.Transactions{}, err
	}

	if body.UserID != userId {
		appErr := util.NewAuthenticationError(errors.New("user ID is not match"))
		return response.Transactions{}, appErr
	}

	input := dto.TransactionsInput{
		UserId:      userId,
		Amount:      body.Amount,
		Description: body.Description,
	}

	err = t.u.ExecTransaction(c.Context(), &input)

	if err != nil {
		return response.Transactions{}, err
	}

	c.SetHeader("Content-Type", "application/json")
	c.SetStatus(http.StatusCreated)

	return response.Transactions{
		Message: "Transaction success",
	}, nil
}
