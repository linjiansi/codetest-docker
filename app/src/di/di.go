package di

import (
	"github.com/jmoiron/sqlx"
	"github.com/linjiansi/codetest-docker/src/handler"
	"github.com/linjiansi/codetest-docker/src/repository"
	"github.com/linjiansi/codetest-docker/src/router/middleware"
	"github.com/linjiansi/codetest-docker/src/usecase"
)

func ProvideAuthenticationMiddleware(db *sqlx.DB) middleware.AuthenticationMiddleware {
	r := repository.NewUserRepository(db)
	u := usecase.NewUserUsecase(r)
	return middleware.NewAuthenticationMiddleware(u)
}

func ProvideTransactionsHandler(db *sqlx.DB) handler.TransactionsHandler {
	r := repository.NewTransactionsRepository(db)
	s := usecase.NewTransactionsUsecase(r)
	return handler.NewTransactionsHandler(s)
}
