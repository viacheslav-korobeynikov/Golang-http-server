package main

import (
	"fmt"
	"net/http"
)

// Функция обработчик
func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Print("Hello")
}

func main() {
	http.HandleFunc("/hello", hello) // Обработка запроса к конкретному запросу
	fmt.Println("Server is listening on port 8081")
	http.ListenAndServe(":8081", nil) // Слушаем TCP-соединение на заданном порту
}
