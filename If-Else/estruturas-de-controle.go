package main

import "fmt"

func main() {
	fmt.Println("Estruturas de Controle") //Condicao if e else
	numero := -5

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual ao 15")
	} //Restrições: if sempre tem que ter as chaves, mesmo sendo uma linha. Também nao pode colocar tudo na mesma linha (fica feio tbm).

	//IF INIT:
	if outronumero := numero; outronumero > 0 { //dividir com ponto e virgula pra avaliar a variavel
		fmt.Println("numero é maior que zero")
	} else if numero < -10 {
		fmt.Println("menor que -10")
	} else {
		fmt.Println("Entre 0 e -10")
	}
	//Quando se cria uma variável usando if init, vc estará limitando ela ao escopo do if!
}
