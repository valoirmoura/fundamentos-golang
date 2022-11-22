package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint16
}

func main() {
	fmt.Println("Print Arquivo Structs")

	var u usuario
	fmt.Println(u) //Imprime com Valor ZERO

	u.nome = "Valoir"
	u.idade = 35
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua Rio Parana", 834}

	//Criação por inferencia
	usuario2 := usuario{"Edicleia", 34, enderecoExemplo}
	fmt.Println(usuario2)

	//Nesse Caso foi inserido apenas o nome no struct
	usuario3 := usuario{nome: "Sidnei"}
	fmt.Println(usuario3)

}
