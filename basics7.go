package main

import "fmt"

type hotdog int // Creating a type to explore atributions and comparisons with created types.

var b hotdog = 101

func main() {
	x := 10
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n\n", b)

	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n\n", b)

	// b = x -> can't do this because they are of different types. Even tho b receives int values, its type is hotdog, not int.

	x = int(b) // now I can convert variable types by declaring it in the atribution
	fmt.Printf("%v\n", x)

}
