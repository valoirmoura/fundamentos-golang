package main

import "fmt"

func main() {

	//[chave]valor
	usuario := map[string]string{
		"nome":      "Valoir",
		"sobrenome": "Moura",
	}
	fmt.Println(usuario)

	//Acessando pela Chave
	fmt.Println(usuario["nome"])

	//Map Aninhado
	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Valoir",
			"segundo":  "Moura",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "Um",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2, "nome") //Deletando a Chave

	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Gemeos",
	}

	fmt.Println(usuario2)

}
