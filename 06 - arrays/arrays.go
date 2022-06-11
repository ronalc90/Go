package main

import "fmt"

func main() {
	var numeros [5]int
	fmt.Println(numeros)

	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30

	fmt.Println(numeros)

	fmt.Println(numeros[3])

	//// arrays con valores

	nombres := [3]string{"ronald", "juan", "pedro"}

	fmt.Println(nombres)

	colores := [...]string{
		"rojo",
		"azul",
		"verde",
	}

	fmt.Println(colores, len(colores))

	/// indices definidos

	monedas := [...]string{
		0: "Dolar",
		2: "Peso",
		4: "Euro"}

	fmt.Println(monedas, len(monedas))

}
