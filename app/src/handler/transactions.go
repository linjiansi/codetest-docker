package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/linjiansi/codetest-docker/src/handler/request"

	"github.com/linjiansi/codetest-docker/src/handler/response"
	"github.com/linjiansi/codetest-docker/src/usecase"
	"github.com/linjiansi/codetest-docker/src/usecase/dto"
	"github.com/linjiansi/codetest-docker/src/util"
)

type TransactionsHandler interface {
	Transactions(w http.ResponseWriter, r *http.Request)
}

type transactionsHandler struct {
	u usecase.TransactionsUsecase
}

func NewTransactionsHandler(u usecase.TransactionsUsecase) TransactionsHandler {
	return &transactionsHandler{u}
}

func (t *transactionsHandler) Transactions(w http.ResponseWriter, r *http.Request) {
	var req *request.Transaction
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.ReturnErrorResponse(w, err)
		return
	}
	ctx := r.Context()
	userId, err := util.GetUserId(ctx)
	if err != nil {
		util.ReturnErrorResponse(w, err)
		return
	}

	if req.UserID != userId {
		appErr := util.NewAuthenticationError(errors.New("user ID is not match"))
		util.ReturnErrorResponse(w, appErr)
		return
	}

	input := dto.TransactionsInput{
		UserId:      req.UserID,
		Amount:      req.Amount,
		Description: req.Description,
	}

	err = t.u.ExecTransaction(ctx, &input)

	if err != nil {
		util.ReturnErrorResponse(w, err)
		return
	}

	res := response.Transactions{
		Message: "Transaction success",
	}

	util.ReturnResponse(w, http.StatusCreated, res)
}
