package main

import (
	"html/template"
	"log"
	"net/http"
)

func index_page(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	tmpl.Execute(w, "Hello World!")

	if err != nil {
		log.Fatal(err)
	}
}

func handleRequest() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index_page)
	log.Println("Запуск сервера на http://127.0.0.1:8000")

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)
}

func main() {
	handleRequest()
}
