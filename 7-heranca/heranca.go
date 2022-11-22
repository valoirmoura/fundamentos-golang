package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint16
	altura    uint16
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {

	p1 := pessoa{"Valoir", "Moura", 35, 172}
	fmt.Println(p1)

	e1 := estudante{p1, "Engenharia", "USP"}
	fmt.Println(e1)
	fmt.Println(e1.nome)

}
