package cart

import (
	"MiayCoffe/bazad" // Пакет, содержащий логику подключения к базе данных
	"MiayCoffe/mnu"
	"os/user"
)

/*
create table carts (
id serial Primary key,
user_id int not null,
products json ,
notes text,
status text ,
created_at text,
total int
);
*/
type Cart struct {
	Id         int         `pg:"id"`
	User_id    int         `pg:"user_id"`
	Products   []CartItems `pg:"products"` //Тут массив описание одного продукта
	Notes      string      `pg:"notes"`
	Status     string      `pg:"status"`
	Created_ad string      `pg:"created_at"`
	Total_sum  int         `pg:"total"`
}

// структура для массива поля Products
type CartItems struct {
	Product  mnu.Product //Структура продукта из пакета Menu
	Quantity int         // Количество товара
}

// func AddToCart(User_id int) error {
// 	db := bazad.GetDB()
// 	var ActiveCart Cart
// 	err := db.Model(&ActiveCart).
// 		Where("user_id = ?", User_id). // Фильтруем по User_id
// 		Where("status = ?", "Active"). // Фильтруем только по активным корзинам
// 		Select()                       // Выполняем выборку из базы данных
// 	if err != nil {
// 		if err.Error() == "pg: no rows in result set" {
// 			var query string = `
// 			INSERT INTO carts (user_id, status)
// 			VALUES (?, ? );`
// 			db.Query(&ActiveCart, query, User_id, "Active")

// 			return nil
// 		}
// 		return err
// 	}
// 	return nil
// }

//Функция для создания корзины 
func AddToCart (user_id) error{ // Она принимает только user_id и возвращает ошибку наверх, чтобы ее можно было обработать.
	db := bazad.GetDB() //объявляем переменную и присваиваем подключение к db
	// Создаем sql запрос SELECT 1- не извлекает значения, а проверят есть ли они, это экономит производительность 
	query := ` 
	SELECT 1
	from carts
	Where user_id = $1 AND status = 'Active'`
	if err != 1 {
		if err.Error() != 1{
			var query string = `INSERT INTO crts (user_id, status) 
			VALUES  `
			
		}
	}





// func Inputcart ()([]Cart,error) {
// 	db := bazad.GetDB()
// 	carts:=[]Cart{}//полученные значения
// 	// z := `INSERT INTO cart(

// 	// )`
// 	// _, err := db.Query(&carts, z)
// 	// if err != nil {
// 	// 	// Если есть ошибка, выводим её в лог с описанием.

// 	// 	return carts, err
// 	// }

// 	db.Model(&carts).Where("category = ?","coffe").Order("category").Select()

// 	_,err:=db.Model(&carts).Insert()
// 	if err != nil {
// 		// Если есть ошибка, выводим её в лог с описанием.

// 		return carts, err
// 	}

// 	for _, product := range carts {
// 		fmt.Println(product)
// 	}
// 	return carts, nil

// } 