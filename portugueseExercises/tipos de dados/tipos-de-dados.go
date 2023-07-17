package main

import (
	"errors"
	"fmt"
)

func main() {
	// Existem 4 tipos de dado para números inteiros no Go: int8, int16, int32 e int64
	// Se você declarar apenas 'int' ele irá usar a arquitetura do seu computador como base - isso acontece quando fazemos inferência de tipos.
	var numero1 int16 = 100
	fmt.Println(numero1)

	// Código abaixo não funciona pois dá overflow, já que o numero não cabe em uma variável que só suporta 8 bits
	//var numero2 int8 = 100000000000000000
	//fmt.Println(numero2)

	// Agora o código funciona pois o valor cabe na variável de 64 bits
	var numero2 int64 = 1000000000000000000
	fmt.Println(numero2)

	// Exemplificando inferência de tipos - sai como 'int'
	numero3 := 200
	fmt.Println(numero3)

	// uint = unsigned int => vale para variáveis sem o sinal definido
	// Código abaixo não funciona pois há a definição do numero negativo
	//var numero4 uint32 = -32
	//fmt.Println(numero4)

	// Agora funciona pois não há a definição
	var numero4 uint32 = 32
	fmt.Println(numero4)

	// alias -> apelidos
	// rune = int32
	// byte = uint8
	var numero5 rune = 12345
	fmt.Println(numero5)

	var numero6 byte = 123
	fmt.Println(numero6)

	// Numeros reais podem ser float32 ou float64
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123456.78
	fmt.Println(numeroReal2)

	// Declaração implícita gera apenas 'float' e segue a sua arquitetura, mas não é possível declarar apenas 'float'!
	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)

	// string -> cadeia de caracteres
	var str string = "abcde"
	fmt.Println(str)

	str2 := "ABCDE"
	fmt.Println(str2)

	// char acontece quando usamos aspas simples -> ele transforma o caractere em um valor numérico pela tabela ascii (só funciona com caracteres da tabela ascii, se não dá erro) e vira tipo int
	char := 'B'
	fmt.Println(char)

	// Valor zero -> valor atribuído se eu não inicializar a variável
	// Cada tipo de dado tem seu valor zero: para string vai ser em branco, pra número será o 0, pra booleano é false, pra erro é <nil>
	var texto string
	fmt.Println(texto)

	var numeroInt int16
	fmt.Println(numeroInt)

	// boleano: true ou false
	var booleano1 bool = true
	var booleano2 bool = false
	fmt.Println(booleano1)
	fmt.Println(booleano2)

	// error: erro é um tipo específico em Go -> valor zero = <nil>
	var erro error
	fmt.Println(erro)

	var erro2 error = errors.New("Erro Interno")
	fmt.Println(erro2)
}
