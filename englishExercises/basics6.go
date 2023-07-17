package main

import "fmt"

func main() {
	x := "Hello!\nHow are you today?\tI hope all is well...\n\n"   // This is an interpreted string. The commands are recognized as commands.
	y := `"Hello!\nHow are you today?\tI hope all is well...\n\n"` // This is a raw string. Every character is recognized as only a character.

	fmt.Println(x)
	fmt.Println(y)

	// fmt package: Print is for printing on the screen, Sprint is for printing on a string (used for variables), and Fprint is for printing on a file or any other writing space (like connections and interfaces)
	// All print statements above follow the same pattern: if followed by 'ln' it will jump a line after printing. if followed by "f" you can format the print as you wish.
}
