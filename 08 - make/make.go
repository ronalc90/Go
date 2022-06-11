package main

import "fmt"

func main() {
	numeros := make([]int, 3, 3)

	numeros[0] = 100
	numeros[1] = 100
	numeros[2] = 100

	numeros = append(numeros, 200)

	fmt.Println(numeros, len(numeros), cap(numeros))
}
