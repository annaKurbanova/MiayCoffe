package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	fmt.Println("hallo, i am alive")

	e := echo.New()
	e.Use(middleware.CORS())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/hello", hello)
	e.GET("/positions", positions)

	e.Logger.Fatal(e.Start(":1323"))
}

func hello(c echo.Context) error {

	data := struct {
		Name  string
		Msg   string
		Phone int
	}{
		Name:  "хмелик",
		Msg:   "HELLO",
		Phone: 123,
	}

	fmt.Println("handle /hello")
	return c.JSON(http.StatusOK, data)
}

func positions(c echo.Context) error {
	str := `
	{
  "position1": {
    "name": "Капучино",
    "price": "200 руб.",
"photo":"https://img.freepik.com/premium-photo/italian-cappuccino-with-decoration_159805-806.jpg?w=826",  
  "category": "coffee"
  },
  "position2": {
    "name": "Латте",
    "price": "220 руб.",
"photo":"https://img.freepik.com/premium-photo/italian-cappuccino-with-decoration_159805-806.jpg?w=826",  
    "category": "coffee"
  },
  "position3": {
    "name": "Эспрессо",
    "price": "100 руб.",
"photo":"https://img.freepik.com/premium-photo/italian-cappuccino-with-decoration_159805-806.jpg?w=826",  
    "category": "coffee"
  },
  "position4": {
    "name": "Раф",
    "price": "300 руб.",
"photo":"https://img.freepik.com/premium-photo/italian-cappuccino-with-decoration_159805-806.jpg?w=826",  
    "category": "coffee"
  },
  "position5": {
    "name": "Ассам",
    "price": "100 руб.",
    "photo": "https://www.chay.info/upload/medialibrary/146/3i03anmkw7cwyvo0v96xi857k1pxqi8a.jpeg",
    "category": "tea"
  },
  "position6": {
    "name": "Сенча",
    "price": "100 руб.",
 "photo": "https://www.chay.info/upload/medialibrary/146/3i03anmkw7cwyvo0v96xi857k1pxqi8a.jpeg",
     "category": "tea"
  },
  "position7": {
    "name": "Горный сбор",
    "price": "100 руб.",
 "photo": "https://www.chay.info/upload/medialibrary/146/3i03anmkw7cwyvo0v96xi857k1pxqi8a.jpeg",
     "category": "tea"
  },
  "position8": {
    "name": "Ягодный",
    "price": "150 руб.",
 "photo": "https://www.chay.info/upload/medialibrary/146/3i03anmkw7cwyvo0v96xi857k1pxqi8a.jpeg",
     "category": "tea"
  },
  "position9": {
    "name": "Облепиха",
    "price": "150 руб.",
 "photo": "https://www.chay.info/upload/medialibrary/146/3i03anmkw7cwyvo0v96xi857k1pxqi8a.jpeg",
     "category": "tea"
  },
  "position10": {
    "name": "Лайм - мята",
    "price": "200 руб.",
    "photo": "https://www.ethnomir.ru/upload/medialibrary/08e/limonad_1140.jpg",
    "category": "limonade"
  },
  "position11": {
    "name": "Малина - базилик",
    "price": "200 руб.",
 "photo": "https://www.ethnomir.ru/upload/medialibrary/08e/limonad_1140.jpg",
     "category": "limonade"
  },
  "position12": {
    "name": "Облепиха - апельсин",
    "price": "200 руб.",
 "photo": "https://www.ethnomir.ru/upload/medialibrary/08e/limonad_1140.jpg", 
    "category": "limonade"
  },
  "position13": {
    "name": "Манго - маракуйа",
    "price": "200 руб.",
 "photo": "https://www.ethnomir.ru/upload/medialibrary/08e/limonad_1140.jpg",
     "category": "limonade"
  },
  "position14": {
    "name": "Чизкейк Сан-Себастьян",
    "price": "250 руб.",
    "photo": "https://www.calend.ru/img/Articles/2019/desert.jpg",
    "category": "deserts"
  },
  "position15": {
    "name": "Тирамису",
    "price": "230 руб.",
 "photo": "https://www.calend.ru/img/Articles/2019/desert.jpg",
     "category": "deserts"
  },
  "position16": {
    "name": "Наполеон",
    "price": "210 руб.",
 "photo": "https://www.calend.ru/img/Articles/2019/desert.jpg",
     "category": "deserts"
  },
  "position17": {
    "name": "Анна Павлова",
    "price": "260 руб.",
 "photo": "https://www.calend.ru/img/Articles/2019/desert.jpg",
     "category": "deserts"
  }
}`

	b := []byte(strings.Replace(strings.Replace(strings.Replace(fmt.Sprintf("%v", str), " ", "", -1), "\n", "", -1), "\t", "", -1))

	fmt.Println("handle /positions")
	return c.JSON(http.StatusOK, string(b))
}
