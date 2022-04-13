package main

import (
	"RestApi/internal/user"
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	log.Println("create router")
	router := httprouter.New()
	log.Println("register user handler")
	handler := user.New()
	handler.Register(router)
	start(router)

}

func start(router *httprouter.Router) {
	log.Println("start application")

	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:           router,
		WriteTimeout:      15 * time.Second,
		ReadHeaderTimeout: 15 * time.Second,
	}

	log.Fatal(server.Serve(listen))
}
