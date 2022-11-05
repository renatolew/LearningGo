package main

import (
	"fmt"
	"strconv"
)

var x int
var y string
var z int

func main() {
	x := 10
	y := "5"
	y_int, err := strconv.Atoi(y)

	if err != nil {
		fmt.Println("Error during conversion")
		return
	}

	z := x + y_int

	fmt.Println(z)

}
