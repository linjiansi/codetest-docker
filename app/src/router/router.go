package router

import "github.com/go-fuego/fuego"

func Run() {
	s := fuego.NewServer(
		fuego.WithAddr(":8888"),
	)

	fuego.Get(s, "/transactions", func(c fuego.ContextNoBody) (string, error) {
		return "Hello, World!", nil
	})

	err := s.Run()
	if err != nil {
		panic(err)
	}
}
