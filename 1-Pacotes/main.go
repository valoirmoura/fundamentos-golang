package main

import (
	"fmt"
	"fundamentos-golang/1-Pacotes/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main!!!")
	auxiliar.Escrever()

	err := checkmail.ValidateFormat("valoir@gmail.com")
	fmt.Println(err)

}
