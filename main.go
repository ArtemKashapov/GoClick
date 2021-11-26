package main

import (
	"html/template"
	"log"
	"net/http"
)

// Создается функция-обработчик "home", которая записывает байтовый слайс, содержащий
// текст "Hello World!" как тело ответа.
func home(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/index.html")
	tmpl.Execute(w, nil)
}

func handleRequest() {
	http.HandleFunc("/", home)

	log.Println("Запуск веб-сервера на http://127.0.0.1:8000")
	err := http.ListenAndServe(":8000", nil)
	log.Fatal(err)
}

func main() {
	handleRequest()
}
