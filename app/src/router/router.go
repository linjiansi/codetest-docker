package router

import (
	"fmt"
	"github.com/go-fuego/fuego"
	"github.com/linjiansi/codetest-docker/src/di"
	"github.com/linjiansi/codetest-docker/src/util"
	"github.com/rs/cors"
)

func Run() {
	db, err := util.NewDB()
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}

	am := di.ProvideAuthenticationMiddleware(db)

	s := fuego.NewServer(
		fuego.WithAddr(":8888"),
		fuego.WithCorsMiddleware(cors.New(cors.Options{
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders: []string{"apikey", "Content-Type"},
		}).Handler),
	)

	transactions := fuego.Group(s, "/transactions")

	fuego.Use(transactions, am.Authentication)

	fuego.Get(transactions, "", func(c fuego.ContextNoBody) (string, error) {
		return "Hello, World!", nil
	})

	err = s.Run()
	if err != nil {
		panic(err)
	}
}
