package main

import "fmt"

// Funcoes Closure: Sao Funções que referenciam variáveis que estão fora so seu corpo
func closure() func() {
	texto := "Dentro da Função Closure!"

	funcao := func() {
		fmt.Println(texto)
	}
	return funcao
}

func main() {

	texto := "Dentro da Funcao Main!"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()

}
