package main

import (
	"fmt"
	"github.com/viacheslav-korobeynikov/Golang-http-server/internal/hello"
	"net/http"
)

func main() {
	router := http.NewServeMux() //Создание собственного роутинга
	hello.NewHelloHandler(router)
	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	} // конфигурация сервера

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe() // Слушаем TCP-соединение на заданном порту
}
