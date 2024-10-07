package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	fmt.Println("helloWork")
	http.HandleFunc("/hello", SeyHihi)
	http.HandleFunc("/miay", handler)
	http.ListenAndServe("127.0.0.1:80", nil)

}

func SeyHihi(w http.ResponseWriter, r *http.Request) {
	fmt.Println("это ми кошки!!!")
	fmt.Fprintf(w, "это ми кошки !!!!!")

}
func handler(w http.ResponseWriter, r *http.Request) {
	// Парсим HTML-шаблон
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Рендерим шаблон в ответе
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
} // подключение html
