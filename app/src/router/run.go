package router

import "net/http"

func Run() {
	server := createServer()
	startServerWithGracefulShutdown(server)
}

func createServer() *http.Server {
	router := configureRoutes()
	return &http.Server{
		Addr:    ":8888",
		Handler: router,
	}
}
