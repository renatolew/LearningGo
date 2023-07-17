package main

import "fmt"

func main() {
	// Printing by importing fmt and calling Println
	fmt.Println("Hello, world!")

	// Variables that are declared have to be called upon, otherwise there will be an error
	_, erros := fmt.Println("Hello, world!", 199, "Hey!")
	fmt.Println(erros)

	// There are three types of variables in go: numbers, strings and bool (boolean)
	x := 10
	y := "string"
	z := true

	fmt.Println(x, y, z)
}
