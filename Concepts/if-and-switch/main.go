package main

import "fmt"

func main() {
	character := "🦹"

	if character == "🦸‍♂️" {
		fmt.Println("Es un super heroe")
	} else if character == "🦹" {
		fmt.Println("Super villain")
	} else {
		fmt.Println("Es un vato ordinario")
	}

	animal := "🦁"
	canSearch := true

	// switch animal {
	// case "🦁":
	// 	fmt.Println("Es un animal salvaje")
	// case "🦛":
	// 	fmt.Println("Es un animal salvaje")
	// case "🐕":
	// 	fmt.Println("Es una mascota")
	// case "🐱":
	// 	fmt.Println("Es una mascota")
	// default:
	// 	fmt.Println("No tenemos ese animal registrado")
	// }

	// switch animal {
	// case "🦁", "🦛":
	// 	fmt.Println("Es un animal salvaje")
	// case "🐕", "🐱":
	// 	fmt.Println("Es una mascota")
	// case "🧑‍🤝‍🧑", "🦸":
	// 	fmt.Println("Es una persona")
	// default:
	// 	fmt.Println("No tenemos ese animal registrado")
	// }
	switch {
	case !canSearch:
		fmt.Println("No esta permitida la busqueda")
	case animal == "🦁" || animal == "🦛":
		fmt.Println("Es un animal salvaje")
	case animal == "🐕" || animal == "🐱":
		fmt.Println("Es una mascota")
	case animal == "🧑‍🤝‍🧑" || animal == "🦸‍♂️":
		fmt.Println("Es una persona")
	default:
		fmt.Println("No tenemos ese animal registrado")
	}
}
