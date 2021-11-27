package main

import (
	_ "github.com/lib/pq"

	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "admin"
	DB_NAME     = "postgres"
)

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
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()

	var lastInsertId int
	err = db.QueryRow("INSERT INTO counter_info(counter) VALUES($1) returning id;", "1").Scan(&lastInsertId)
	checkErr(err)
	fmt.Println("last inserted id =", lastInsertId)

	handleRequest()
}

func checkErr(err error) {

	if err != nil {
		panic(err)
	}
}
