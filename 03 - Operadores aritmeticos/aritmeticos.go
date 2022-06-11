package main

import "fmt"

func main() {
	a := 20
	b := 10

	result := a + b

	fmt.Println("Suma", result)

	result = a - b
	fmt.Println("Resta", result)

	result = a * b

	fmt.Println("Multiplicacion", result)

	result = a / b

	fmt.Println("Division", result)

	result = a % b

	fmt.Println("Modulo", result)

	result = a << b

	fmt.Println("Shift left", result)
}
