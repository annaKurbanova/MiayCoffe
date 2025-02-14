package mnu

import (
	"MiayCoffe/bazad" // Пакет, содержащий логику подключения к базе данных
	"fmt"             // Пакет для вывода на экран
	// Пакет для логирования ошибок
)

// Структура для представления данных о продукте
// Каждое поле структуры соответствует колонке таблицы в базе данных.
type Product struct {
	Id          int    // Уникальный идентификатор продукта (соответствует колонке Id в таблице)
	Name        string // Название продукта (соответствует колонке Name в таблице)
	Description string // Описание продукта (соответствует колонке Description в таблице)
	Price       int    // Цена продукта (соответствует колонке Price в таблице)
	Category    string // Категория продукта (соответствует колонке Category в таблице)
	Picture     string // Ссылка на изображение продукта (соответствует колонке Picture в таблице)
}

type ProductForCart struct{
	Product
	Quantity int
}

// Функция для получения списка продуктов из базы данных
func GetProduct() ([]Product, error) {
	// Создаем срез (массив) структуры Product, который будет заполнен данными из базы
	Tovar := []Product{}

	// Подключение к базе данных:
	// `bazad.GetDB()` возвращает ссылку на объект подключения, который был инициализирован в пакете bazad.
	db := bazad.GetDB() // в этой переменной хранится подключение к бд

	z := `SELECT * FROM products` //запрос в бд
	//эта функция возвращает 2 значения, поэтому нам нужны 2 переменные для хранения. ошибку игнорировать нельзя, поэтому мы объявляем переменную err, а вторая нам не нужна(поэтому прочерк). и не забывать про амперсант &!
	_, err := db.Query(&Tovar, z) // через Query мы указываем то куда сохранить (Tovar), и что сделать(z).
	// Проверяем, возникла ли ошибка при выполнении запроса.
	if err != nil {
		// Если есть ошибка, выводим её в лог с описанием.

		return Tovar, err

	}
	//запустить
	for _, product := range Tovar {
		fmt.Printf("ID: %d, Name: %s, Description: %s, Price: %d, Category: %s, Picture: %s\n",
			product.Id, product.Name, product.Description, product.Price, product.Category, product.Picture)
	}
	return Tovar, nil
}
