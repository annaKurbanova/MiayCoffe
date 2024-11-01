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
	// удалить
	fs := http.FileServer(http.Dir("./static"))

	// Обрабатываем запросы к "/static/" и перенаправляем на файловый сервер
	fmt.Println("helloWork")
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/hello", SeyHihi)
	http.HandleFunc("/miay1", handler)
	http.HandleFunc("/registration", hendlerGuest)
	http.HandleFunc("/miay2", PodkluchenieHTMLsCC)
	//http://127.0.0.1:1324/miay2
	//Ручка для входа
	http.HandleFunc("/Vhod", KeisDliaVxoda)
	//регистрация пользователя
	http.HandleFunc("/Registration", KeisDliaAvtorizacii)
	http.ListenAndServe(":1324", nil)
	//http.ListenAndServe("127.0.0.1:1324", nil)

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
	avtorization.CreateUser(Person.Name, Person.PhoneNumber)
}

// это тип для запроса, какие данные нам должны вернуться
type DataGuest struct {
	Name        string `json:"Name"`
	PhoneNumber string `json:"PhoneNumber"`
}

func KeisDliaAvtorizacii(w http.ResponseWriter, r *http.Request) {
	var p DataGuest

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	avtorization.Registration(p.Name, p.PhoneNumber)
	fmt.Println("Сервер запущен на http://localhost:1324")
	w.WriteHeader(http.StatusOK)
}

func KeisDliaVxoda(w http.ResponseWriter, r *http.Request) {
	var p DataGuest
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	avtorization.Vxod(p.Name, p.PhoneNumber)
	fmt.Println("Сервер запущен на http://localhost:1324")
	w.WriteHeader(http.StatusOK)
}
