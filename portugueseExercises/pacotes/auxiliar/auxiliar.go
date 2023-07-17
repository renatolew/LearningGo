package auxiliar

import "fmt"

// o que criarmos que começar com letra maiuscula é entendido como público e pode ser exportado/importado por outros pacotes
// o que criarmos que começar com letra minuscula é entendido como privado e somente funciona dentro do pacote em que for escrito
func Escrever() {
	fmt.Println("Escrevendo do pacote auxiliar")
	escrever2() // funcionou pois agora faz parte de uma função declarada de forma pública
}
