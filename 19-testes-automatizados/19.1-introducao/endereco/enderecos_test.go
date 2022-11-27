package enderecos

import (
	"strings"
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	//Começa obrigatóriamente com a Palavra Test com T maiusculo
	//O Nome do Arquivo Obrigatóriamente deve ter _test no final...

	t.Parallel() //Roda os Testes em PAraralelo

	cenariosDeTestes := []cenarioDeTeste{
		{"Rua Rio Parana", "rua"},
		{"Avenida Brasil", "avenida"},
		{"Alameda Dom Pedro I", "tipo invalido!"},
		{"Estrada da Graciosa", "estrada"},
		{"Rodovia dos Minérios", "rodovia"},
		{"RUA MARACAI", "rua"},
	}

	for _, cenario := range cenariosDeTestes {
		tipoEnderecoRecebido := strings.ToLower(TipoDeEndereco(cenario.enderecoInserido))

		if tipoEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido é diferente do esperado! Esperava-se %s e recebeu %s", tipoEnderecoRecebido, cenario.retornoEsperado)
		}

	}
	//
	//enderecoParaTeste := "rua Paulista"
	//tipoDeEnderecoEsperado := "rua"
	//tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)
	//
	//if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
	//	t.Errorf("O tipo recebido é diferente do esperado! Esperava-se %s e recebeu %s", tipoDeEnderecoEsperado, tipoDeEnderecoRecebido)
	//}

}

//Comando go test ./... Testa todos os Diretórios...
//Comando go test -v //Verboso tras mais dados no Test
//Comando go test --cover //Retorna % de cobertura do Testes
//Coamndo go test --coverprofile enderecos_cobertura //retorna funções que não está Cobertas
//Comando go tool cover --func=enderecos_cobertura testa o arquivo txt (enderecos_cobertura) gerando do comando acima
//Comando go tool cover --html=enderecos_cobertura //Retorna um Arquivo .html contendo todas as funçoes que não foram cobertas e o que não foi coberto... O MELHOR DOS COMANDOS.... (obs. precisa ter o .txt)...

func TestQualquerCoisa(t *testing.T) {

	t.Parallel() //Roda os Testes em Paraleleo

	if 1 > 2 {
		t.Error("Teste quebrou")
	}
}
