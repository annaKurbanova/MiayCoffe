package cart

import (
	"MiayCoffe/mnu"
	"MiayCoffe/bazad" // Пакет, содержащий логику подключения к базе данных
)


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
	Quantity int          // Количество товара
}

func AddToCart(User_id int, Products int, Notes string, quantity int, Total_sum int) error {
	db := bazad.GetDB()
	var ActiveCart Cart
	err := db.Model(&ActiveCart).
		Where("user_id = ?", User_id). // Фильтруем по User_id
		Where("status = ?", "Active"). // Фильтруем только по активным корзинам
		Select()                       // Выполняем выборку из базы данных
	if err != nil {
		if err.Error() == "pg: no rows in result set" {
			// todo sql запрос создание новой строки с корзиной!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
		}
		return err  
	}
	return nil
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
