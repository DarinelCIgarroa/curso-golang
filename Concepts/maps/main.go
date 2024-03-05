package main

import "fmt"

func main() {
	// Definir un map
	music := make(map[string]string)
	music["guitar"] = "🎸"
	music["violin"] = "🎻"
	fmt.Println(music)
	// otra forma de definir un map
	things := map[string]string{
		"star":  "✨",
		"mouse": "🖱️",
	}
	things["phone"] = "📱"
	delete(things, "star") // Elimnar un elemento

	fmt.Println(things)
	content, ok := things["mouse"]
	fmt.Println("things", things)
	fmt.Println("content", content, ok)
}
