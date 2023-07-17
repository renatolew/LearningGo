package main

import "fmt"

var z = 100

func main() {
	y := 50
	printingAnyNumber(y)
	printingAnyNumber(z)
	//printingAnyNumber(a)  -- it doesn't work because 'a' is not declared as var outside the scope neither declared inside the code block

	otherFunc()
}

func otherFunc() {
	a := 200

	fmt.Println(a, "is used here!")
}

func printingAnyNumber(x int) {
	fmt.Println(x)
}
