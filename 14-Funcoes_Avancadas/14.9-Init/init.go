package main

import "fmt"

// A Função INIT é executada antes de tudo, inclusive da função Main, e pode ser utilizada uma por arquivo, não uma por pacote
func init() {
	fmt.Println("Executando a Função Init!!!")
}

func main() {

	fmt.Println("Executando a Função Main!")
}
