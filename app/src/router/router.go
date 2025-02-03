package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/linjiansi/codetest-docker/src/di"
	"github.com/linjiansi/codetest-docker/src/util"
	"github.com/rs/cors"
)

func configureRoutes() http.Handler {
	db, err := util.NewDB()
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}

	am := di.ProvideAuthenticationMiddleware(db)
	th := di.ProvideTransactionsHandler(db)

	r := chi.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"apikey", "Content-Type"},
	})
	r.Use(c.Handler)

	r.Route("/transactions", func(r chi.Router) {
		r.Use(am.Authentication)
		r.Post("/", th.Transactions)
	})

	return r
}
