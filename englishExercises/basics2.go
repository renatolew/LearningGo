package main

import "fmt"

// You can declare a variable outside of the main scope and it becomes avaliable everywhere. Declaring variables outside of scopes is done by using =, not :=
var a = "Hey!"

func main() {
	// The := is for declaring, while = is only for setting a value to a variable.
	x := 10.0
	y := "Good morning!"
	fmt.Printf("x: %v, %T", x, x)
	fmt.Printf("\ny: %v, %T\n\n", y, y)

	// If you are going to declare at least one variable, you can use :=
	x, z := 20, 50
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("z: %v, %T\n\n", z, z)

	// If you are only setting new values to variables, you need to use =
	x = 15
	fmt.Printf("x: %v, %T\n\n", x, x)
	
	// Variables set and declared outside of this scope can be used within it, but variables set and declared inside the scope can be used only within it
	fmt.Printf("a: %v, %T", a, a)