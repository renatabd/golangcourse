package main

import "fmt"

func closure() func() {
	texto := "Dentro da funcao closure"

	funcao := func() {
		fmt.Println(texto)
	}
	return funcao
}

func main() { //Funcoes que referenciam variaveis que estao fora do seu corpo
	texto := "Dentro da funcao main"
	fmt.Println(texto)

	funcaoNova := closure() //Imprimiu o texto da funcao closure
	funcaoNova()
}
