package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

// DE STRUCT PARA JSON
func main() {
	c := cachorro{"Rex", "Dalmata", 3}
	fmt.Println(c)

	//O Marshal transforma um Struct em JSON
	cachorroEmJSON, erro := json.Marshal(c)
	if erro != nil {
		log.Fatalln(erro)
	}

	fmt.Println(cachorroEmJSON)

	//O bytes.NewBuffer converte o bytes em JSON
	fmt.Println(bytes.NewBuffer(cachorroEmJSON))

	c2 := map[string]string{
		"nome": "Toby",
		"raca": "Poodle",
	}

	cachorro2EmJSON, err := json.Marshal(c2)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(cachorro2EmJSON)
	fmt.Println(bytes.NewBuffer(cachorro2EmJSON))
}
