package main

import (
	"fmt"
)

func main() {

	// keyExistance()

}

// Проверьте наличие ключа в map.
// если обращаешься по не существующему ключу выведи "не найденно"
func keyExistance() {
	mapa := map[string]string{
		"one": "один",
		"two": "два",
	}

	fmt.Println("mapa[one]: ", mapa["one"])

	fmt.Println("mapa[three]", mapa["three"])

	fmt.Println("не найденно")
}
