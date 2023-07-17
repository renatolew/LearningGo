package main

// Importando o pacote auxiliar
import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

// Executando a função criada no pacote auxiliar
func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()
	// auxiliar.escrever2() - Não funcionou pois está declarado de forma privada em outro pacote

	erro := checkmail.ValidateFormat("devbook@gmail.com")
	fmt.Println(erro) // Devolve nulo pois o formato é válido

	erro2 := checkmail.ValidateFormat("123")
	fmt.Println(erro2) // Devolve erro de formato inválido
}

// o arquivo modulo.exe precisa ser atualizado sempre que fizermos mudanças, e para isso é preciso rodar novamente o comando 'go build' no terminal, na pasta do pacote.
// se você rodar 'go install' em vez de 'go build' o pacote ficará disponível na pasta raiz do Go no computador, em vez de apenas na pasta do projeto.
// para importar pacotes externos, utiliza-se o comando 'go get endereco-do-pacote' - ao realizar essa operação o arquivo go.mod é atualizado automaticamente com a dependência
// o comando 'go mod tidy' remove todas as dependências que não estão sendo utilizadas do arquivo go.mod
