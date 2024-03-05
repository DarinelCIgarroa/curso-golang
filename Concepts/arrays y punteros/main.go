package main

import "fmt"

func main() {
	// Definir el array
	things := [...]string{"🫎", "🦁", "🐢", "🍌", "🍎", "🍊", "🍓"}
	fmt.Println("things ->", things)
	// Mofificar la posicion de un array
	things[0] = "🐅"
	fmt.Println("things ->", things)
	/*
		SLICES
		Un slice es una varible que crea punteros de un array
		se puede definir un slice a partir de un array el primer parametro
		es incluyente y el segundo excluyente
	*/
	// Definir un slice
	animals := things[0:2]
	// Otra manera de definir un slice a partir de un array
	animals2 := things[4:]
	fmt.Println("animals slice ->", animals)
	fmt.Println("animals2 slice ->", animals2)
	fruit := things[3:7]
	fmt.Println("fruit slice->", fruit)
	animals[0] = "🐞"
	fruit[3] = "🥥"
	fmt.Println("-----------------")
	fmt.Println("things array ->", things)
	fmt.Println("animals slice ->", animals)
	fmt.Println("fruit slice ->", fruit)
	// Definir un slice vacio
	carModels := []string{}
	carModels = append(carModels, "Jetta")  // Agregar un elemento al slice
	carModels = append(carModels, "Camaro") // saber la capacidad tamano de carModels
	fmt.Println(len(carModels))             // saber el valor tamano de carModels
	fmt.Println(cap(carModels))             // saber la capacidad tamano de modelsCars
	fmt.Println("car models", carModels)

}
