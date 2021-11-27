package main

import (
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	tmpl.Execute(w, "Hello World!")

	if err != nil {
		log.Fatal(err)
	}
}

func handleRequest() {
	http.HandleFunc("/", home)
	log.Println("Запуск сервера на http://127.0.0.1:8000")
	err := http.ListenAndServe(":8000", nil)
	log.Fatal(err)
}

func main() {
	handleRequest()
}
