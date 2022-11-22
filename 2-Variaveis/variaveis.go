package main

import "fmt"

func main() {

	//Variável String: Forma 1 Tipo Explicito
	var variavel1 string = "Variável 1"

	//Variavel String: Forma 2 Tipo Inferido
	variavel2 := "Variável 2"

	//Declaração de Variável em Bloco
	var (
		variavel3 string = "Variável 3"
		variavel4 string = "Variável 4"
	)

	//Declaração de Variável em Bloco por Inferencia de Tipo
	variavel5, variavel6 := "Variável 5", "Variável 6"

	//Constante: Forma 1 Tipo Explicito
	const constante1 string = "Constante 1"

	fmt.Println(variavel1)
	fmt.Println(variavel2)
	fmt.Println(variavel3)
	fmt.Println(variavel4)
	fmt.Println(variavel5)
	fmt.Println(variavel6)
	fmt.Println(constante1)

	//Invertendo Valores das Variáveis
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5)
	fmt.Println(variavel6)

}
