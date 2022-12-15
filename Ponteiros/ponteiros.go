package main

import "fmt"

//Vai salvar o endereço de memória de algo
func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++ //Não muda a variavel2. Quando se atribui um valor a uma variavel, está se criando uma cópia. Da linha 10 pra baixo, as duas variaveis perdem a ligação uma com a outra.
	fmt.Println(variavel1, variavel2)

	//PONTEIRO É UMA REFERÊNCIA DE MEMÓRIA
	//Está passando o endereço do local onde ta armazenado o valor de uma variavel. Ou seja, ele não cria uma cópia do valor, ele utiliza o que está naquele endereço...

	var variavel3 int //valor inteiro
	var ponteiro *int //endereço de memória de um valor inteiro

	variavel3 = 100
	ponteiro = &variavel3 //& comercial para indicar um ponteiro

	//Como ver o valor que está salvo no endereço de memória: colocando um asterisco na frente de ponteiro - desreferenciacao

	fmt.Println(variavel3, ponteiro)

	variavel3 = 777
	fmt.Println(variavel3, *ponteiro)
}
