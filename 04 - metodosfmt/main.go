package main

import "fmt"

func main() {
	hola := "hola"
	mundo := "mundo"

	fmt.Println(hola)
	fmt.Println(mundo)
	nombre := "ronald"
	edad := 20

	fmt.Printf("hola, %s, tienes %d años \n", nombre, edad)

	fmt.Printf("hola, %v, tienes %v años \n", nombre, edad)

	mensaje := fmt.Sprintf("hola, %v, tienes %v años \n", nombre, edad)

	fmt.Println(mensaje)

	fmt.Scanln(&nombre)

	fmt.Println("otro nombre :", nombre)
}
