package enderecos

import (
	"golang.org/x/exp/slices"
	"strings"
)

func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}
	enderecoEmLetraMinuscula := strings.ToLower(endereco)
	primeiraPalavraDoEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0]

	enderecoTemUmTipoValido := slices.Contains(tiposValidos, primeiraPalavraDoEndereco)

	//Ou sem implementação
	//enderecoTemUmTipoValido2 := false
	//for _, tipo := range tiposValidos {
	//	enderecoTemUmTipoValido = true
	//}

	if enderecoTemUmTipoValido {
		return primeiraPalavraDoEndereco
	}

	return "Tipo Invalido!"

}
