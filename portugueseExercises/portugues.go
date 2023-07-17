// Arquivos .go sempre começam chamando o pacote main
package main

// Para importar qualquer coisa se usa a função import e coloca-se entre parenteses cada biblioteca, função ou pacote importado - linha por linha
import (
	"fmt"
)

// Se inicia o código a ser executado chamando a função main, que precisa ter o mesmo nome do pacote inicializado no inicio do arquivo (package main)
func main() {
	fmt.Println("Olá, mundo!")
}
