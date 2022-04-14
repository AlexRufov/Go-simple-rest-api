package main

import (
	"RestApi/internal/user"
	"RestApi/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"time"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("create router")
	router := httprouter.New()
	logger.Info("register user handler")
	handler := user.New(logger)
	handler.Register(router)
	start(router)

}

func start(router *httprouter.Router) {
	logger := logging.GetLogger()

	logger.Info("start application")

	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:           router,
		WriteTimeout:      15 * time.Second,
		ReadHeaderTimeout: 15 * time.Second,
	}

	logger.Fatal(server.Serve(listen))
}
