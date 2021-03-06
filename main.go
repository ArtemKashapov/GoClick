package main

import (
	"encoding/xml"
	"io/ioutil"
	"strconv"

	_ "github.com/lib/pq"

	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "admin"
	DB_NAME     = "postgres"
)

func handleRequest() {
	mux := http.NewServeMux()

	// Обрабатываем запросы
	mux.HandleFunc("/", index_page)
	mux.HandleFunc("/result", result_page)
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

func result_page(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/result.html")
	doc2XML()
	tmpl.Execute(w, nil)

	if err != nil {
		log.Fatal(err)
	}
}

func click_handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		countValue := insertValue()
		fmt.Fprintf(w, strconv.Itoa(countValue))
	}
}

func main() {
	handleRequest()
}

func doc2XML() {
	db := dbConnection()
	defer db.Close()

	var value int
	err := db.QueryRow("SELECT counter from counter_info WHERE id = $1;", 1).Scan(&value)
	checkErr(err)

	type Clicks struct {
		XMLName xml.Name `xml:"clicker_count"`
		Clicks 	int `xml:"clicks"`
	}

	v := &Clicks{Clicks: value}

	output, err := xml.MarshalIndent(v, "  ", "    ")
	checkErr(err)

	// os.Stdout.Write(output)
	_ = ioutil.WriteFile("data/click_count.xml", output, 0644)
}

func insertValue() (countValue int) {
	db := dbConnection()

	var value int
	err := db.QueryRow("SELECT counter from counter_info WHERE id = $1;", 1).Scan(&value)
	checkErr(err)

	err = db.QueryRow("UPDATE counter_info SET counter=$1 WHERE id = $2 returning counter;", value+1, 1).Scan(&countValue)
	checkErr(err)
	// fmt.Println("last inserted id =", countValue)
	defer db.Close()

	return
}

func dbConnection() (db *sql.DB) {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	return db
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
