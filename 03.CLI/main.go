package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {

	fmt.Println("Rudiney Iniciando CLI...")

	//'app' é o nome do pacote/package e Gerar é uma função pública
	aplicacao := app.Gerar()

	//paramertro 'os.Args' é padrão para aplicação de linha de comando
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}

	//Para executar:
	//Testando local
	//-> go run main.go servidores --host pismo.io
	//-> go run main.go ip --host pismo.io

	//Teste com arquivo executável
	//Fazer o build da aplicação para gerar o executável
	//-> go build
	////este comando criará um executável com o mesmo nome do módulo, definido em go.mod: linha-de-comando
	//Executar o aplicação
	//-> ./linha-de-comando ip --host pismo.io

}
