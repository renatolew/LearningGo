package main

import "fmt"

// Definindo uma função
func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// Declarando da forma abaixo, todos os parametros terão tipo int8
// Declarando da forma abaixo, podemos ter mais de um retorno por função - Isso ajuda depois a lidar com erros e exceções
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

// É sempre preciso abrir a função main para chamar outras funções. A função main é a que executará seu código, e em um pacote 'main' não tem como rodar nada sem uma função 'main'
func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	// funções também são um tipo, então é possível declarar uma variável como função e chamá-la depois, executando a função.
	var f = func() {
		fmt.Println("Função F")
	}
	f()

	// Outra forma de declarar e utilizar uma função é deixar o parametro em aberto, definindo apenas o que será utilizado como entrada, e definindo esse parâmetro ao chamar a função
	// Dessa forma o tipo da função muda um pouco, pois o tipo do parâmetro faz parte do tipo da função
	var f2 = func(txt string) {
		fmt.Println(txt)
	}
	f2("Texto da função 2")

	// Dessa forma definimos também o retorno, e precisamos então declarar esse retorno na função. O código abaixo, então, printa a string duas vezes, pois uma vez ela é printada na execução da função, e na segunda vez pelo comando de printar o retorno.
	var f3 = func(txt string) string {
		fmt.Println(txt)
		return txt
	}
	resultado := f3("Texto da função 3")
	fmt.Println(resultado)

	// O código abaixo não funciona pois a função de calculo retorna dois valores, e temos apenas uma variável recebendo esses valores.
	//resultadoCalculos := calculosMatematicos(10, 15)
	//fmt.Println(resultadoCalculos)

	// Agora funciona pois declaramos duas variáveis para receber esses valores
	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	// Ao utilizar o 'underline' (_) estamos dizendo pro Go ignorar o resultado da segunda operação. Dessa forma conseguimos ter os dois retornos da função de cálculos e utilizar apenas um
	resultadoSoma2, _ := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma2)

	// A ordem importa muito aqui, pois define qual dos retornos será ignorado.
	_, resultadoSubtracao2 := calculosMatematicos(10, 15)
	fmt.Println(resultadoSubtracao2)
}
