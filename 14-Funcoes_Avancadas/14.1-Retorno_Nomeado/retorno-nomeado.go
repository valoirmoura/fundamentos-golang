package main

import "fmt"

// Quando colocamos nome no Retorno, não se faz necessário adicionar no return o nome das variáveis, pois já foram declaradas na função
func calculosMatematicos(n1, n2 int) (soma, subracao int) {
	soma = n1 + n2
	subracao = n1 - n2
	return
}

func main() {
	soma, subtracao := calculosMatematicos(10, 20)
	fmt.Println(soma, subtracao)
}
