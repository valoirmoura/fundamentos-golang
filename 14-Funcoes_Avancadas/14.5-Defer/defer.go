package main

import "fmt"

//O DEFER em Ingles significa adiar, ou seja, a função anotada com defer será a última a ser executada

func funcao1() {
	fmt.Println("Executando a funcao1!!!")
}

func funcao2() {
	fmt.Println("Executando a Funcao2 !!!")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Media Calculada o Resultado será retornado!")
	fmt.Println("Entrando na Função para Verificar se o Aluno está Aprovado!")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}
	return false
}

func main() {

	defer funcao1()
	funcao2()

	fmt.Println(alunoEstaAprovado(2, 8))

}
