package main

/*
INSTALANDO MODULOS EXTERNOS
	EXEMPLO:
		go get github.com/badoux/checkmail

	O comando acima vai atualizar o arquivo go.mod,
		escrevendo a nova dependência nele.
	E para utilizar, faz-se como a seguir

REMOVENDO MODULOS EXTERNOS NÃO UTILIZADOS
	go mod tidy
		Este comando vai apagar o import da
			dependência que existe no arquivo go.mod
*/
import (
	"fmt"
	"moduloTeste/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")

	/*
		Função começando com letra minuscula
		é função privada do PACOTE.
		Por isso escrever2 não está
		disponivel aqui
		Por isso Escrever está disponivel
	*/
	//auxiliar.escrever2()

	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("rudiney1h@gmail.com")
	erro2 := checkmail.ValidateFormat("rudiney1hgmail.com")
	fmt.Println(erro)
	fmt.Println(erro2)

}
