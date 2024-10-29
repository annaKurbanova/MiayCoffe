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
	http.HandleFunc("/miay1", handler)
	http.HandleFunc("/registration", hendlerGuest)
	http.HandleFunc("/miay2", PodkluchenieHTMLsCC)
	//http://127.0.0.1:1324/miay2
	//регистрация пользователя
	http.HandleFunc("/Registration", KeisDliaAvtorizacii)
	http.ListenAndServe("127.0.0.1:1324", nil)

}
func PodkluchenieHTMLsCC(w http.ResponseWriter, r *http.Request) {
	// Парсим HTML-шаблон
	tmpl, err := template.ParseFiles("index2.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Рендерим шаблон в ответе
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func SeyHihi(w http.ResponseWriter, r *http.Request) {
	fmt.Println("это ми кошки!!!")
	fmt.Fprintf(w, "это ми кошки !!!!!")

}
func handler(w http.ResponseWriter, r *http.Request) {
	// Парсим HTML-шаблон
	tmpl, err := template.ParseFiles("index1.html")
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

func KeisDliaAvtorizacii(w http.ResponseWriter, r *http.Request) {
	fmt.Println("KeisDliaAvtorizacii")
	type Person struct {
		Login       string
		Password    string
		PhoneNumber int
	}
	var p Person

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	avtorization.Registration(p.Login, p.Password, p.PhoneNumber)

}
