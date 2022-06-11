package main

import "fmt"

func main() {
	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(numeros, len(numeros))

	numeros = append(numeros, 0)
	numeros = append(numeros, 11)
	fmt.Println(numeros, len(numeros))

	subSlicen := numeros[3:6]

	numeros[4] = 100

	fmt.Println(subSlicen)
	fmt.Println(numeros)

	//Slicen ... puntero
	// capacidad
	// longitud

	meses := []string{"Enero", "Febrero", "Marzo"}

	fmt.Printf("Len : %v, cap: %v, %p \n", len(meses), cap(meses), meses)

	meses = append(meses, "Abril")

	fmt.Printf("Len : %v, cap: %v, %p \n", len(meses), cap(meses), meses)

}
