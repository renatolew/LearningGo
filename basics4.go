package main

import "fmt"

// variable types in go are static
var x int = 10
var z string

// z = "hello" doesn't work because if you declare a variable without setting a value, you can only do that inside a code block

func main() {

	//x = "hey" - doesnt work because x was declared as int, not string
	//x = 20.5 - doesnt work because x was declared as int, not float
	x = 20
	z = "hello"
	fmt.Println(x)
	fmt.Println(z)
}
