package main

import (
	"fmt"
	enderecos "fundamentos-golang/19-testes-automatizados/19.1-introducao/endereco"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("Rua Paulista")
	fmt.Println(tipoEndereco)
}
