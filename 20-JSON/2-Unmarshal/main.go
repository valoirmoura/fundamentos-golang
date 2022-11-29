package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"-"` //Quando deixado um - essa sera ignorado na conversao para JSON
}

// DE JSON PARA STRUCT
func main() {

	cachorroEmJSON := `{"nome":"Rex","raca":"Dalmata","idade":3}`
	fmt.Println(cachorroEmJSON)

	//Necessita Criar uma Variável para ser compreendido
	var c cachorro

	//Neste caso especifico como estamos criando na unha, a variável cachorroEmJSON é string, e o UnMarshal necessita de Slice de Bytes
	cachorroEmSliceDeBytes := []byte(cachorroEmJSON)

	//Passamos o Slice de Bytes, e em seguinda Apontamos para a Variável que queremos que estes dados sejam inseridos
	if err := json.Unmarshal(cachorroEmSliceDeBytes, &c); err != nil {
		log.Fatalln(err)
	}

	fmt.Println(c)

	cachorro2EmJSON := `{"nome" : "Toby", "raca":"Poodle"}`
	c2 := make(map[string]string)

	cachorro2EmJSONBytes := []byte(cachorro2EmJSON)

	if erro := json.Unmarshal(cachorro2EmJSONBytes, &c2); erro != nil {
		log.Fatalln(erro)
	}

	fmt.Println(c2)

}
