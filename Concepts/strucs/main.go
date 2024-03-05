package main

import "fmt"

type Person struct {
	Name        string
	Age         uint
	HasChildren bool
}

func main() {
	darinel := Person{
		Name:        "Darinel",
		Age:         26,
		HasChildren: false,
	}
	darwin := &darinel
	darwin.Age = 25
	fmt.Printf("%+v\n", darinel)
	fmt.Printf("%+v\n", darwin)
}
