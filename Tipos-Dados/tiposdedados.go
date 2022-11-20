package main

import (
	"errors"
	"fmt"
)

func main() {
	//1- NUMEROS INTEIROS
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
	//FIM NUMEROS INTEIROS

	//2- NUMEROS REAIS
	//numeros quebrados
	//float32, float64
	var numeroReal1 float32 = 4.55
	fmt.Println(numeroReal1)

	//Inferencia de tipos:
	numeroReal2 := 12345.55
	fmt.Println(numeroReal2)
	//FIM NUMEROS REAIS

	//3- STRINGS
	var str string = "Texto String"
	fmt.Println(str)

	//String Inferencia de tipos:
	str2 := "Texto inferencia de string"
	fmt.Println(str2)

	//Char (Ele vai ser um Int/Rune)
	//Ele declara o número do B na tabela ASCII
	char := 'B'
	fmt.Println(char)
	//FIM STRINGS

	//VALOR ZERO
	//Valor que vai ser atribuido a uma variavel quando não se indica um valor

	var texto string
	fmt.Println(texto)
	//Retorna uma linha em branco

	var number int16
	fmt.Println(number)
	//Retorna um numero zero

	//BOOLEANO E ERRO
	var booleano1 bool = true
	fmt.Println(booleano1)

	var booleano2 bool
	fmt.Println(booleano2)
	//Vai retornar o valor zero = false

	var erro error
	fmt.Println(erro)
	//valor zero dele é <nil>. é um tipo de dado que serve como valor zero pra muitas coisas (pra interface, erros, ponteiros, etc)

	var erro2 error = errors.New("Erro interno")
	fmt.Println(erro2)
	//Importou um package lá em cima em import
	//O tipo erro é MUITO usado em Go
	//FIM BOOLEANO E ERRO
}
