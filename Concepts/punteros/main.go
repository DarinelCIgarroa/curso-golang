package main

import "fmt"

func main() {
	color := "🟥"
	pointerColor := &color
	*pointerColor = "🟦"
	fmt.Printf("tipo -> %T valor -> %v direccion -> %v \n", color, color, &color)
	fmt.Printf("tipo: %T, valor: %v, desreferenciacion: %s \n", pointerColor, pointerColor, *pointerColor)
}
