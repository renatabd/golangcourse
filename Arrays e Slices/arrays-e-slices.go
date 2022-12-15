package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	//ARRAYS - Lista de valores
	var array1 [5]int //todos os dados tem de ser o mesmo tipo
	array1[0] = 30    // sempre começar por 0
	fmt.Println(array1)

	array2 := [5]string{"P1", "P2", "P3", "P4", "P5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5} //Fixa o numero de arrays de acordo com a qtd de itens que coloquei depois das chaves em int. Não é muito dinamico.
	fmt.Println(array3)

	//SLICE - baseada no array com flexibilidades, princ. com tamanho
	//Muito usado em GO! NÃO É UM ARRAY. Parece, mas nao é.
	slice := []int{10, 12, 17, 29, 32, 77} //ARRAY INTERNO, sempre se referencia a algo
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 18) //adiciona um valor a mais
	fmt.Println(slice)

	slice2 := array2[1:3] //indice 1 a 3, equivalente a maior ou igual a 1 e menor que 3 (inclusivo, exclusivo)
	fmt.Println(slice2)

	array2[1] = "Posição alterada"
	fmt.Println(slice2)

	//Arrays Internos
	//Funcao make aloca um espaço na memória para determinada coisa
	slice3 := make([]float32, 10, 11) //parametros dentro dos parametros tipo, tamanho, capacidade
	fmt.Println(slice3)

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)

	fmt.Println(slice3)
	fmt.Println(len(slice3)) //Imprimir tamanho (lenght)
	fmt.Println(cap(slice3)) //Capacidade

	slice4 := make([]float32, 5)
	slice4 = append(slice4, 10)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}
