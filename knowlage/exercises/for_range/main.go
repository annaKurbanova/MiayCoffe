package main

import (
	"fmt"
	"math/rand"
)

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
	fmt.Println("отсортированные по убыванию числа: ")
	var masItog []int //3,5,1
	for i := 0; i < len(numbers); i++ {
		var index int
		var sum int
		index, sum = min(numbers) // 1
		numbers[index] = -1
		masItog = append(masItog, sum)

	}
	fmt.Println(masItog)

	// Используем функцию Max чтобы получить не 1 число а массив.
	var masMaksItog []int // Объявили массив для записи возрастающей сортировки.
	for i := 0; i <= len(numbers); i++ {
		var index int // переменные для хранения индекса и значения.
		var sum int
		index, sum = max(numbers)
		numbers[index] = -1
		masMaksItog = append(masMaksItog, sum)
	}
	fmt.Println("отсортированные по возрастанию числа:", (masMaksItog))

}

func min(nums []int) (int, int) {
	var result int = 100
	var j int
	for i, value := range nums {

		if value < 0 {
			continue
		}
		if value < result {
			result = value
			j = i
		}

	}

	return j, result
}
func max(nums []int) (int, int) { // создаем массив и прописываем входные данные(nums []int) и выходные данны (int, int). Два инта потому что первый это индекс, второй значения как в range
	var resultMax int = 0 // объявляем переменную котрая будет хранить 1 максимальное число
	var jM int            // переменная которую мы преровняем к i чтобы в forRange было видно наш i(индекс)
	for i, value := range nums {
		if value < 0 {
			continue
		}
		if value > resultMax {
			resultMax = value
			jM = i
		}
	}
	return jM, resultMax

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
