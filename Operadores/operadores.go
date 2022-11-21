package main

import "fmt"

func main() {
	// ARITMETICOS
	//+, -, /, *, % (MOD/remainder)
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 2
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 3

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	//EXEMPLO:
	//var n1 int16 = 10
	//var n2 int32 = 25
	//soma := n1 + n2
	//fmt.println(soma)
	//O exemplo acima não vai funcionar pois não se pode fazer nada no Go com variaveis que sao de tipos diferentes

	var n1 int16 = 10
	var n2 int16 = 25
	soma2 := n1 + n2
	fmt.Println(soma2)

	//FIM ARITMÉTICOS

	//ATRIBUIÇÃO
	var variavel1 string = "string"
	variavel2 := "string2"
	fmt.Println(variavel1, variavel2)
	//FIM ATRIBUIÇÃO

	//OPERADORES RELACIONAIS (retornam apenas resultados de true/false)
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2) //!= significa diferente
}
