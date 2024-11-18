package bazad

import (
	"context"
	"log"

	"github.com/go-pg/pg/v10" // Пакет для работы с базой данных PostgreSQL
)

// Структура клиента
type Client struct {
	tableName struct{} `sql:"Client" pg:"Client"`
	ID          int64 `sql:"ID" pg:"ID"`
	Name        string `sql:"Name" pg:"Name"`
	PhoneNumber string `sql:"Phone_Number" pg:"Phone_Number"`
}

var db *pg.DB

// Подключение к базе данных
func Connect() {
	db = pg.Connect(&pg.Options{
		Addr:     "localhost:5432", // Адрес вашей базы данных
		User:     "postgres",       // Имя пользователя
		Password: "superuser",      // Ваш пароль
		Database: "miay",           // Имя базы данных
	})

	// Проверка подключения
	err := db.Ping(context.Background())
	if err != nil {
		log.Fatalf("Ошибка при подключении к базе данных: %v", err)
	}
}

// Получаем подключение к базе данных
func GetDB() *pg.DB {
	return db
}
