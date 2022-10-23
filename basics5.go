package main

import "fmt"

var a int     // zero value is 0
var b float64 // zero value is 0.0
var c string  // zero value is ""
var d bool    // zero value is false

func main() {
	fmt.Printf("a: %v, %T\n", a, a)
	fmt.Printf("b: %v, %T\n", b, b)
	fmt.Printf("c: %v, %T\n", c, c)
	fmt.Printf("d: %v, %T\n", d, d)
}
