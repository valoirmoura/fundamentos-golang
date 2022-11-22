package main

import "fmt"

//Passar qualquer coisa para a função independente do tipo de dados...
func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("String!!!")
	generica(1)
	generica(true)

	fmt.Println()
}
