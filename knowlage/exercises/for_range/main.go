package main

import (
	"fmt"
	"math/rand"
	
)

type Sortirovka []int

// Реализация метода Len для Sortirovka
func (a Sortirovka) Len() int {
	return len(a)
}

// Реализация метода Swap для Sortirovka
func (a Sortirovka) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// Реализация метода Less для Sortirovka
func (a Sortirovka) Less(i, j int) bool {
	return a[i] < a[j]
}
func main() {
	forRange()

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
	ItogMass := []int{}
	for _, value := range numbers {
		if value%2 == 0 {
			ItogMass = append(ItogMass, value)

		}

	}

	fmt.Printf("numbers: %v\n", numbers)

	// вывести тут четные числа
	fmt.Println("четные числа: ", ItogMass)
	//sort.Ints(numbers)
	// вывести тут отсортированные числа
	fmt.Println("отсортированные числа: ", numbers)

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
