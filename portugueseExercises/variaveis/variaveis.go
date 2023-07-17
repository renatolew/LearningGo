package main

import "fmt"

func main() {
	// Declarando variável de forma explicita, dizendo que é uma variável e definindo o tipo dela
	var variavel1 string = "Variável 1"
	fmt.Println(variavel1)

	// Declarando variável de forma implicita, apenas atribuindo o valor e deixando que o Go defina o tipo por conta própria a partir da declaração (inferência de tipo)
	variavel2 := "Variável 2"
	fmt.Println(variavel2)

	// Também é possível declarar desta forma:
	var (
		variavel3 string = "Variável 3"
	)
	fmt.Println(variavel3)

	// Declarando mais de uma variável ao mesmo tempo
	var (
		variavel4 string = "Variável 4"
		variavel5 string = "Variável 5"
	)
	fmt.Println(variavel4, variavel5)

	// Constantes seguem a mesma lógica, a unica diferença é que não é possível alterar o valor de uma constante.
	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	// Com Go é possível intercambiar o valor de duas variáveis sem precisar de uma variável auxiliar!
	variavel2, variavel1 = variavel1, variavel2
	fmt.Println(variavel1, variavel2)
}
