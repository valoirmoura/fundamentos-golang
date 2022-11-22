package main

import (
	"fmt"
)

func main() {

	//Em GO a única forma de loop é FOR

	//Tipo While em GO
	//i := 0
	//for i < 10 {
	//	i++
	//	fmt.Println("Incrementando i")
	//	time.Sleep(time.Second)
	//}
	//fmt.Println(i)

	//Forma 2 de For
	//for j := 0; j < 10; j += 2 {
	//	fmt.Println("Incrementando j ", j)
	//	time.Sleep(time.Second)
	//}

	//Forma 3
	//nomes := []string{"Valoir", "Edicleia", "Nerli", "Sidnei"}
	//for indice, nome := range nomes {
	//	fmt.Println(indice, nome)
	//}

	//Iterando String: As Letras que ele traz, são os Códigos da Tabela ASC
	//for indice, letra := range "VALOIR" {
	//	fmt.Println(indice, letra)
	//}

	//para trazer as letras precisa utilizar a biblioteca String
	//for indice, letra := range "VALOIR" {
	//	fmt.Println(indice, string(letra))
	//}

	//For sobre o MAP
	usuario := map[string]string{
		"nome":      "Valoir",
		"sobrenome": "Moura",
	}
	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	//Loop Infinito
	//for {
	//	fmt.Println("Executando Loop Infinito!!!")
	//	time.Sleep(time.Second)
	//}

}
