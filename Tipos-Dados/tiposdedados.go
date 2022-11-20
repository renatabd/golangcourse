package main

import "fmt"

func main() {
	//NUMEROS INTEIROS
	//4 tipos (Tambem suportam numeros negativos)
	//int8, int16, int32, int64
	//Se eu utilizasse um numero tipo 100000000 no int8 - ele daria uma mensagem de overflow
	//O int sozinho usa a arquitetura do computador como base

	var numero int16 = 100
	fmt.Println(numero)

	//Inferencia de tipos:

	numero1 := -100000000000
	fmt.Println(numero1)

	//unsigned int - int sem sinal
	var numero2 uint32 = 10000
	fmt.Println(numero2)

	//alias: rune = int32 (usando em caracteres utilizando a tabela ascii)
	var numero3 rune = 123456
	fmt.Println(numero3)

	//alias: byte = uint8
	var numero4 byte = 7
	fmt.Println(numero4)
}
