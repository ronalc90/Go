package main

import "fmt"

func main() {
	dias := make(map[int]string)

	// agregando datos
	dias[1] = "domingo"
	dias[2] = "lunes"
	dias[3] = "martes"
	dias[4] = "miercoles"
	dias[5] = "jueves"
	dias[6] = "viernes"
	dias[7] = "sabado"

	fmt.Println(dias)

	delete(dias, 1)

	fmt.Println(dias)
}
