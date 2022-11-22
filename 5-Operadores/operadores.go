package main

import "fmt"

func main() {
	//Aritméticos (+, -, /, *, %)
	somar := 1 + 2
	subtracao := 2 - 1
	divisao := 10 / 5
	multiplicacao := 5 * 10
	modulo := 10 % 2

	fmt.Println(somar, subtracao, divisao, multiplicacao, modulo)

	var numero1 int16 = 10
	var numero2 int32 = 25
	//soma := numero1 + numero2 //Não Funciona, pois são tipos diferentes um é int32 a outra é int16
	fmt.Println(numero2, numero1)

	//Atribuição
	var variavel1 string = "String" //Atribuição Explicita
	variavel2 := "String"           //Atribuição implicita
	fmt.Println(variavel1, variavel2)

	//Relacionais (Retornam Bool)
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	//Lógicos (Só Existem 3, não existem o OU Exclusivo aquelas bagaça chata....
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso) //Operador E
	fmt.Println(verdadeiro || falso) //Operador OU
	fmt.Println(!verdadeiro)         //Operador Negacão

	//Unários (Em GO só existe pós-fixado, não existe pré-fixado igual no JAVA)
	numero := 10
	numero++ //numero + 1
	fmt.Println(numero)
	numero += 5 //(numero = numero + 5)
	fmt.Println(numero)
	numero *= 3 //(numero = numero * 3)
	fmt.Println(numero)
	numero %= 3 //numero = numero % 3
	fmt.Println(numero)

	//Operador TERNARIO: Bah!!! Não existe no GO.... somente o IF e ELSE tradicional
	//A Premissa do GO é que haja apenas uma maneira de fazer as coisas, com exceção a Declaração de VARIÁVEIS, por isso, não existe o meu amado TERNÁRIO!!!!

}
