package main

import "fmt"

func main() {

	var variavel1 int = 10
	var variavel2 int = variavel1

	variavel1++
	fmt.Println(variavel1, variavel2)

	//Ponteiro é uma Referência de memória
	var variavel3 int
	var ponteiro *int

	fmt.Println(variavel3, ponteiro)

	variavel3 = 100
	ponteiro = &variavel3
	fmt.Println(variavel3, *ponteiro) //desreferenciacao

	variavel3 = 150
	fmt.Println(variavel3, *ponteiro)

}
