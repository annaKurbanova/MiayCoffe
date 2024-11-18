package avtorization

import (
	"fmt"
	"log"
	"MiayCoffe/bazad" // Подключаем пакет для работы с базой данных
)

// Создаем нового пользователя в базе данных
func CreateUser(Name string, PNumber string) {
	// Создаем объект клиента
	client := &bazad.Client{
		Name:        Name,
		PhoneNumber: PNumber,
	}

	// Получаем объект базы данных и сохраняем нового клиента
	db := bazad.GetDB()
	_, err := db.Model(client).Insert()
	if err != nil {
		log.Printf("Ошибка при создании пользователя: %v", err)
	} else {
		fmt.Println("Пользователь успешно зарегистрирован")
	}
}

// Проверка входа пользователя
func Vxod(Name string, PNumber string) bool {
	// Ищем пользователя в базе данных по имени и номеру телефона
	client := &bazad.Client{}
	db := bazad.GetDB()
	err := db.Model(client).Where("name = ? AND phone_number = ?", Name, PNumber).Select()

	if err != nil {
		fmt.Println("Пользователь не найден")
		return false
	}

	fmt.Println("Успешный вход")
	return true
}