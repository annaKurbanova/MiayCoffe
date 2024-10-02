package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// forRange()

	// matrix()

}

/*
вывести четные числа из слайса numbers используя range.
отсортировать по возрастанию numbers используя range.
*/
func forRange() {
	numbers := []int{}
	for i := 1; i < 10; i++ {
		numbers = append(numbers, rand.Intn(10))
	}

	fmt.Printf("numbers: %v\n", numbers)

	// вывести тут четные числа
	fmt.Println("четные числа: ")

	// вывести тут отсортированные числа
	fmt.Println("отсортированные числа: ")
}

/*
Есть матрица с цифрами.
просуммировать элементы на главной диагонали.
Пример:
+-----+
|1 2 3|
|4 5 6|
|7 8 9|
+-----+
должно получиться: 1 + 5 + 9 = 15
*/
func matrix() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Printf("matrix: %v\n", matrix)

	var sum int = 0
	// тут тебе поможет вложенный цикл

	fmt.Println("сумма элементов главной диагонали: ", sum) // должно получиться 15
}
