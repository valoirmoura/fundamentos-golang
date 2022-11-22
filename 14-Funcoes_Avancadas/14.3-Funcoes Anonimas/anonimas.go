package main

import "fmt"

func main() {

	//Função Anonima, (sem nome), no final do bloco da função abre e fecha parentêses, isso significa que assim que acabarmos de declarar a função já será executada, não precisa chama-la
	func() {
		fmt.Println("Olá Mundo!!!")
	}()

	//Passando Parâmetro
	func(texto string) {
		fmt.Println(texto)
	}("Olá Mundo!!! Passando Parâmetro!!!")

	//Função Anonima com Retorno
	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %v", texto)
	}("Passando Parâmetro com Retorno!!!")
	fmt.Println(retorno)

}
