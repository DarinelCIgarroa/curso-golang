package main

import "fmt"

func main() {

	number := 5
	// Manera clasica de recorrer un for
	for i := 0; i < number; i++ {
		fmt.Println("indice", i)
	}
	// Segunda menera de recorre un for
	e := 1
	for e < number {
		fmt.Println("indice", e)
		e++
	}
	// Tercera manera de recorre un for
	u := 0
	for {
		if u == 5 {
			break
		}
		u++
		fmt.Println(u)
	}
	// Recorrer un array
	array := [...]int{1, 2, 3, 4, 5}
	for index, value := range array {
		fmt.Println("index", index, "value", value)
	}
	// Recorrer un slice
	things := []string{"📷", "📦", "🔗", "📱"}
	for index, value := range things {
		fmt.Println("index", index, "value", value)
	}
	// Recorrer un map
	animals := map[string]string{
		"lion":   "🦁",
		"turtle": "🐢",
		"horse":  "🐴",
		"tiger":  "🐅",
		"shark":  "🦈",
	}
	// Remplazar un elemento del map
	for key, value := range animals {
		if key == "horse" {
			animals[key] = "🐎"
		}
		fmt.Println(key, value)
	}
	fmt.Println(animals)
	// Recorrer un string
	myName := "Darinel"

	for index, value := range myName {
		fmt.Printf("index %v, value %v \n", index, string(value))
		fmt.Printf("%T tipo \n", value)
	}
}
