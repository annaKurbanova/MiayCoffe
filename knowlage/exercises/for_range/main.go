package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//forRange()
	matrix2()

}

/*
вывести четные числа из слайса numbers используя range.
отсортировать по возрастанию numbers используя range.
*/

func forRange() {

	numbers := []int{}
	numbers2 := []int{}

	for i := 0; i < 10; i++ {
		// r:=rand.Intn(10)
		// numbers = append(numbers, r)
		// numbers2 = append(numbers2, r)

		numbers = append(numbers, rand.Intn(10))
		numbers2 = append(numbers2, numbers[i])
	}
	ItogMass := []int{}
	for _, value := range numbers {
		if value%2 == 0 {
			ItogMass = append(ItogMass, value)

		}

	}

	fmt.Printf("numbers: %v\n", numbers)
	fmt.Printf("numbers2: %v\n", numbers2)

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

	fmt.Printf("numbers после: %v\n", numbers) // -1
	fmt.Printf("numbers2 после: %v\n", numbers2) // рандом

	// Используем функцию Max чтобы получить не 1 число а массив.
	var masMaksItog []int // Объявили массив для записи возрастающей сортировки.
	for i := 0; i <= len(numbers2); i++ {
		var index int // переменные для хранения индекса и значения.
		var sum int
		index, sum = max(numbers2)
		numbers2[index] = -1
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
	fmt.Printf("matrix: %vggg\n",5)
	fmt.Println(matrix[0])
	fmt.Println(matrix[1])
	fmt.Println(matrix[2])
	var su int
	j :=0
	for i := 0; i < len(matrix); i++ {
		su= su+matrix[i][j]
		j++
	}


	
	// тут тебе поможет вложенный цикл

	fmt.Println("сумма элементов главной диагонали: ", su) // должно получиться 15
}
/* вывести сумму четных чисел матрицы = 20
Пример:
+-----+
|1 2 3|
|4 5 6|
|7 8 9|
+-----+
должно получиться: 2 + 4 + 6 + 8 = 20 */
func matrix2 (){
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Printf("matrix: %vggg\n",5)
	fmt.Println(matrix[0])
	fmt.Println(matrix[1])
	fmt.Println(matrix[2])
	var su int
	
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++{
			if matrix[i][j]%2==0{
				su= su+ matrix[i][j]
			}

		}

	}
	fmt.Println("сумма элементов главной диагонали: ", su) // должно получиться 15
}