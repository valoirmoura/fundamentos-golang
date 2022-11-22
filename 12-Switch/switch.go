package main

import "fmt"

func diaDaSemana(numero int) string {
	//Forma 1
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sabado"
	default:
		return "Dia Inválido!"
	}
}

func diaDaSemana2(numero int) string {
	var diaDaSemana string

	//Forma 2
	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
		fallthrough //joga o resultado para a variável dia da semana e vai para a proxima verificação, ou seja, ele guarda o valor do resuldado na variável, e vai para a proxima condição com esse novo valor
	case numero == 2:
		diaDaSemana = "Segunda-Feira"
		fallthrough
	case numero == 3:
		diaDaSemana = "Terça-Feira"
	case numero == 4:
		diaDaSemana = "Quarta-Feira"
	case numero == 5:
		diaDaSemana = "Quinta-Feira"
	case numero == 6:
		diaDaSemana = "Sexta-Feira"
	case numero == 7:
		diaDaSemana = "Sabado"
	default:
		diaDaSemana = "Dia Inválido!"
	}

	return diaDaSemana
}

func main() {

	dia := diaDaSemana(5)
	dia2 := diaDaSemana2(10)

	fmt.Println(dia)
	fmt.Println(dia2)
}
