package auxiliar

import "fmt"

//"Comentario obrigatorio para funções públicas"
func Escrever() {
	fmt.Println("Escrevendo do pacote auxiliar função Escrever")

	/*
		Por se encontrar no mesmo pacote,
		a função que é PRIVADA DO PACOTE
		poderá ser executada
	*/
	escrever2()
}
