package main

import "fmt"

// Funcoes Recursivas, são funcoes que dependem delas mesmas, ou seja, o resultado é o parâmetro da proxima... e assim por diante. Funcoes Recursivas, necessitam informar o ponto de parada... se não Estouro de Pilha...
// FIBONACCI [1 1 2 3 5 8 13 21 ....] sempre soma o anterior...
// Evitar usar Funções Recursivas quando existem muitos dados, pois ela é lenta.....
func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {

	posicao := uint(21)
	for i := uint(0); i < posicao; i++ {
		fmt.Println(fibonacci(i))
	}
}
