package main

import "fmt"

// O Panic mata o sistema, ou seja, ele para o Sistema não é um Erro!, entretando antes de ele parar ele chama todas as funções DEFER
func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A Média é Exatamente 6!!!")
}

// Recover tenta recuperar um PANIC
func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execução Recuperada Com Sucesso!")
	}
}

func main() {
	fmt.Println(alunoEstaAprovado(6, 6))
	fmt.Println("Pós Execução!")
}
