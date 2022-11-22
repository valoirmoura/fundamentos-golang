package auxiliar

import "fmt"

func Escrever() {
	fmt.Println("Escrevendo do Arquivo do Pacote Auxiliar!!!")

	//Usando a função não exportável dentro do mesmo pacote
	escrever2()
}
