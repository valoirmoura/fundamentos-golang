package main

import (
	"errors"
	"fmt"
)

func main() {

	//Tipos Inteiros int8, int16, int32, int64 (o int sozinho utiliza a Arquitetura do Comutador como base, ou seja, se o Computador for 32bits utiliza 32 se for 64 utiliza 64)
	var numero int16 = 100
	fmt.Println(numero)

	//Tipos Inteiros uint8, uint16, uint32, uint64 (suporta apenas numeros positivos)
	var numero2 uint64 = 200
	fmt.Println(numero2)

	//alias (O rune é um apelido ao int32)
	var numero3 rune = 1234
	fmt.Println(numero3)

	//alias (o byte é um apelido ao uint8)
	var numero4 byte = 123
	fmt.Println(numero4)

	//Tipos Reais São somente 2: float32 e float64
	var numeroReal1 float32 = 123.34
	var numeroReal2 float64 = 123.23
	numeroReal3 := 123.34 //Por inferencia (A arquitetura é conforme o seu computador se for 32 é float32 se for 64 é float64)
	fmt.Println(numeroReal1)
	fmt.Println(numeroReal2)
	fmt.Println(numeroReal3)

	//String
	var str1 string = "texto"
	str2 := "Texto 2" //por inferencia
	fmt.Println(str1)
	fmt.Println(str2)

	//O GO não possui o Tipo Char
	char := 'B' //ISSO não é um Caractere e sim o Codigo na Tabela ASC
	fmt.Println(char)

	//Tipo BOOLEAN
	var bol bool = true
	fmt.Println(bol)

	//Tipo ERRO (para criar um Erro é necessário utilizar o Pacote errors do GO)
	var err error = errors.New("erro Aqui")
	fmt.Println(err)

	//VALOR ZERO DAS VARIÁVEIS

	var texto string //Valor Zero é em Branco sou seja nada
	fmt.Println(texto)

	var inteiro int16 //Valor Zero do Tipo Int é 0 (Zero)
	fmt.Println(inteiro)

	var float float64 //Valor Zero para Tipo Float é 0 (Zero)
	fmt.Println(float)

	var bol1 bool //Valor Zero para Booleano é FALSE
	fmt.Println(bol1)

	var err1 error //Valor Zero de ERROR é NIL
	fmt.Println(err1)

}
