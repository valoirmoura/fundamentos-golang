package main

import (
	"fmt"
	"reflect"
)

func main() {

	var array1 [5]string
	array1[0] = "Posicao 1"
	fmt.Println(array1)

	//Array possui um tamanho fixo e não pode ser alterado
	array2 := [5]string{"Posicao 1", "Posicao 2", "Posicao3", "Posicao 4", "Posicao 5"}
	fmt.Println(array2)

	//Nesse Caso o Tamanho do Array será com a quantidade de itens que foi informado.
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	//No Slice não é especificado o Tamanho, ele é variável, também é tipado
	slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	//Adicionando valor ao Slice
	slice = append(slice, 7)
	fmt.Println(slice)

	//O Slice2 contem as informações do indice 1 ao 2 do array2, ele ignora o 3
	slice2 := array2[1:3]
	fmt.Println(slice2)

	//O Slice funciona como um ponteiro de Array
	array2[1] = "Posicao Alterada!"
	fmt.Println(slice2)

	//Arrays Internos
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice3 = append(slice3, 5)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice3 = append(slice3, 6)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

}
