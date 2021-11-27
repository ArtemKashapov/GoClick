package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func handleRequest() {
	mux := http.NewServeMux()

	// Обрабатываем запросы
	mux.HandleFunc("/", index_page)
	mux.HandleFunc("/click", click_handler)

	// Информирование
	log.Println("Запуск сервера на http://127.0.0.1:8000")

	// Подключаем стили и скрипты
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Слушаем порт 8000:
	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)
}

func index_page(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")

	tmpl.Execute(w, nil)

	if err != nil {
		log.Fatal(err)
	}
}

func click_handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		
		fmt.Println("Receive ajax post data string ", json.NewDecoder(r.Body))
		w.Write([]byte("Done!"))
	}
}

func main() {
	handleRequest()
}
