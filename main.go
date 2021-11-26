package main

import (
	"log"
	"net/http"
)

func index_page(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func handleRequest() {
	http.HandleFunc("/", index_page)

	log.Println("Запуск веб-сервера на http://127.0.0.1:8000")
	err := http.ListenAndServe(":8000", nil)
	log.Fatal(err)
}

func main() {
	handleRequest()
}
