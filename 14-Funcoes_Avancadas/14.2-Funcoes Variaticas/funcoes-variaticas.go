package main

import "fmt"

// Funções para receber quantos numeros quiser, muito Show né
func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

// Apenas é possível um parâmetro variático por função, e também o parâmetro variático necessita ser exclusivamente o último da função...
func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	totalDaSoma := soma(1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println(totalDaSoma)

	escrever("Olá Mundo!!!", 10, 20, 30, 40, 50, 60)
}
