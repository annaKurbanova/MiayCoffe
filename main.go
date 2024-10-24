package main

import (
	"MiayCoffe/knowlage/servis/avtorization"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	// Указываем путь к папке с файлами
	fs := http.FileServer(http.Dir("./static"))

	// Обрабатываем запросы к "/static/" и перенаправляем на файловый сервер
	fmt.Println("helloWork")
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/hello", SeyHihi)
	http.HandleFunc("/miay", handler)
	http.HandleFunc("/registration", hendlerGuest)
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

// добавить новую обработку(ручку)для создания пользователя
func hendlerGuest(w http.ResponseWriter, r *http.Request) {
	var Person DataGuest
	err := json.NewDecoder(r.Body).Decode(&Person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
    avtorization.CreateUser(Person.Name, Person.Password, Person.PhoneNumber)
}

type DataGuest struct {
	Name        string
	Password    string
	PhoneNumber int
}
