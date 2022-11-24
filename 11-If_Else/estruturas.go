package main

import "fmt"

func main() {

	numero := -5

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	//IF INIT
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Numero é Maior que Zero!")
	} else if numero < -10 {
		fmt.Println("É menor do que -10")
	} else {
		fmt.Println("Entre 0 e -10")
	}

}
