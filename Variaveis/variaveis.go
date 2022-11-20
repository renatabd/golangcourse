package main

import "fmt"

func main() {
	// declaracao de variaveis
	var variavel1 string = "Variavel 1"

	// Inferência de tipos: variavel string oculta: (ele determina o tipo da variavel pelo valor dela, como sao letras entao é string)
	variavel2 := "Variavel 2"
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	//variaveis explicitas
	var (
		variavel3 string = "variavel3"
		variavel4 string = "variavel4"
	)

	fmt.Println(variavel3, variavel4)

	// Inferência de tipos: para mais de uma variavel
	variavel5, variavel6 := "variavel5", "variavel6"
	fmt.Println(variavel5, variavel6)

	//Constante de variavel: nao muda o valor
	const constante1 string = "constante1"
	fmt.Println(constante1)

	//Inverter o valor das variaveis
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}
