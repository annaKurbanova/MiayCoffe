package main

//Верните несколько значений функции.
import (
	"fmt"
	"sort"
)

func main() {
	delenie, ostDel := myFunction(5, 6)
	fmt.Println(delenie, ostDel)
	x := ludi() //объявили переменную в которую будет передаваться значение из функции ludi
	fmt.Println(x)
	z :=Slaisi()
	fmt.Println(z)

}
//-------------------------------------------------------------------------------------------------------
//Верните несколько значений функции.
func myFunction(z int, x int) (int, int) {

	delenie := z / x
	ostDel := z % x
	return delenie, ostDel

}

// 1. Входные данные:  a = 10, b = 3
//         |
// 2. Функция Divide:
//     - result = a / b  ->  result = 10 / 3 = 3
//     - remainder = a % b  ->  remainder = 10 % 3 = 1
//         |
// 3. Возвращаем:
//     - result = 3
//     - remainder = 1
//         |
// 4. Сохраняем значения в переменные:
//     - res = 3
//     - rem = 1
//         |
// 5. Используем res и rem в коде (выводим на экран)

// -----------------------------------------------------------------------------------------
// Инициализируйте структуру.
type person struct {
	age  int
	name string
	rost int
}

func ludi() person {

	Anna := person{23, "Anna", 163}
	return Anna
}

// Схематическое представление:
// Определение структуры:
// Person
// ├── Name (string)
// ├── Age (int)
// └── Height (float64)
// Инициализация структуры:
// person := Person{
//     Name: "Alice",
//     Age: 30,
//     Height: 1.70,
// }
// Структура person теперь содержит следующие данные:
// person
// ├── Name: "Alice"
// ├── Age: 30
// └── Height: 1.70
// Доступ к полям:
// person.Name → "Alice"
// person.Age → 30
// person.Height → 1.70

// --------------------------------------------------------------------------------------------------------------
// Отсортируйте элементы слайса целых чисел.
func Slaisi() []int {
	slais := []int{1, 12, 22, 5, 7, 23, 55}
	sort.Ints(slais)
	return slais
}
//Схематическое представление:
// Исходный слайс:
// [5, 3, 8, 1, 2, 7]
// Вызов функции сортировки:
// sort.Ints(numbers)
// Отсортированный слайс:
// [1, 2, 3, 5, 7, 8]