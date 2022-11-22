package main

import "fmt"

// Funcao somar recebe 2 parametros e retorna um int8
func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// Funções em GO podem ter mais de um Retorno
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	//Funções em GO são cidadãos de primeira Classe, ou seja, o valor de uma Variável pode ser o retorno de uma função
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) {
		fmt.Println(txt)
	}
	f("Texto da Função 1!")

	//Utilizando a Função com Multiplos Retornos
	resultadoSoma, resultadoSubracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubracao)

	//Neste caso o "_" significa que Ignora a Variável, pois no GO é obrigatório utilizar todas as Variáveis
	resultadoSoma2, _ := calculosMatematicos(10, 20)
	fmt.Println(resultadoSoma2)

}
