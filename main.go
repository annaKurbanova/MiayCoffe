package main

import (

	// Пакет для работы с авторизацией
	"MiayCoffe/bazad" // Пакет для работы с базой данных
	avtorization "MiayCoffe/knowlage/servis"
	"MiayCoffe/mnu"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {

	// Подключаемся к базе данных
	bazad.Connect()

	fmt.Println("Сервер работает!")

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/hello", SeyHihi)
	http.HandleFunc("/miay1", handler)
	http.HandleFunc("/miay2", PodkluchenieHTMLsCC)
	http.HandleFunc("/registration", CorsTest)
	http.HandleFunc("/Vhod", KeisDliaVxoda)
	http.HandleFunc("/Menu", KeisMenu)

	fmt.Println("Сервер запущен на http://localhost:1324")
	http.ListenAndServe(":1324", nil)
}

func SeyHihi(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Тестовый вывод в консоль")
	fmt.Fprintf(w, "Это Ми Кошки!")
}

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index1.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func PodkluchenieHTMLsCC(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index2.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func CorsTest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		fmt.Println("Обработан preflight-запрос")
		w.WriteHeader(http.StatusOK)
		return
	}

	var person DataGuest
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	avtorization.CreateUser(person.Name, person.PhoneNumber)

	w.WriteHeader(http.StatusOK)
}

type DataGuest struct {
	Name        string `json:"Name"`
	PhoneNumber string `json:"PhoneNumber"`
}

func KeisDliaVxoda(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	var person DataGuest
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	avtorization.Vxod(person.Name, person.PhoneNumber)
	fmt.Println("Запрос на вход пользователя получен")

	w.WriteHeader(http.StatusOK)
}
func KeisMenu(w http.ResponseWriter, r *http.Request) {
	fmt.Println("выполняется KeisMenu")
	//это корсы, это надо просто надо
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		fmt.Println("Обработан preflight-запрос")
		w.WriteHeader(http.StatusOK)
		return
	}
	tovars, err := mnu.GetProduct()
	if err != nil {
		log.Printf("Ошибка при выполнении запроса: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = json.NewEncoder(w).Encode(tovars)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
