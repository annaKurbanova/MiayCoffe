package main

import (
	"MiayCoffe/knowlage/servis/avtorization" // Импортируем модуль для работы с авторизацией
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

// Основная функция, запускающая сервер
func main() {
	// Создаем файловый сервер для обслуживания статических файлов (CSS, JS и т.п.)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs)) // Маршрутизация для статических файлов

	// Маршрутизация для обработки различных URL-адресов
	http.HandleFunc("/hello", SeyHihi)                  // Маршрут для теста
	http.HandleFunc("/miay1", handler)                  // Маршрут для первого HTML
	http.HandleFunc("/miay2", PodkluchenieHTMLsCC)      // Маршрут для второго HTML
	http.HandleFunc("/registration", CorsTest)          // Маршрут для регистрации с поддержкой CORS
	http.HandleFunc("/Vhod", KeisDliaVxoda)             // Маршрут для входа пользователя

	fmt.Println("Сервер запущен на http://localhost:1324")
	http.ListenAndServe(":1324", nil) // Запускаем сервер на порту 1324
}

// Функция CorsTest обрабатывает запросы с поддержкой CORS для маршрута /registration
func CorsTest(w http.ResponseWriter, r *http.Request) {
	// Устанавливаем заголовки для разрешения CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")        // Разрешает доступ с любого источника
	w.Header().Set("Access-Control-Allow-Methods", "POST")    // Разрешает метод POST
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type") // Разрешает заголовок Content-Type

	// Обрабатываем preflight-запросы (OPTIONS)
	if r.Method == http.MethodOptions {
		fmt.Println("Обработан preflight-запрос")
		w.WriteHeader(http.StatusOK)
		return
	}

	// Обработка основного POST-запроса
	var person DataGuest
	err := json.NewDecoder(r.Body).Decode(&person) // Декодируем тело запроса в структуру DataGuest
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Вызываем функцию создания пользователя из модуля авторизации
	avtorization.CreateUser(person.Name, person.PhoneNumber)

	// Успешный ответ
	w.WriteHeader(http.StatusOK)
}

// Подключение второго HTML-шаблона к маршруту /miay2
func PodkluchenieHTMLsCC(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index2.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil) // Рендерим шаблон в ответ
}

// Функция для тестового маршрута /hello
func SeyHihi(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Тестовый вывод в консоль")
	fmt.Fprintf(w, "Это Ми Кошки!") // Ответ в браузер
}

// Подключение первого HTML-шаблона к маршруту /miay1
func handler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index1.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil) // Рендерим шаблон в ответ
}

// Структура данных пользователя для обмена данными с клиентом
type DataGuest struct {
	Name        string `json:"Name"`
	PhoneNumber string `json:"PhoneNumber"`
}

// Обработчик для входа пользователя
func KeisDliaVxoda(w http.ResponseWriter, r *http.Request) {
	// Проверяем, что метод запроса — POST
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	// Декодируем данные пользователя из запроса
	var person DataGuest
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Вызываем функцию входа пользователя из модуля авторизации
	avtorization.Vxod(person.Name, person.PhoneNumber)
	fmt.Println("Запрос на вход пользователя получен")

	// Успешный ответ
	w.WriteHeader(http.StatusOK)
}
