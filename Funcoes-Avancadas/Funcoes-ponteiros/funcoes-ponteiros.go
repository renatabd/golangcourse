package main

import "fmt"

func inverterSinal(numero int) int {
	return numero * -1 //Passando parametro por copia
}

func inverterSinalComPonteiro(numero *int) {
	*numero = *numero * -1 //Passando uma referencia para a funcao
	//Saber diferencas e impactos que passar uma variavel com e sem ponteiro pois pode mudar de acordo com o pacote
}

func main() {
	numero := 20
	numeroInvertido := inverterSinal((numero))
	fmt.Println(numeroInvertido)
	fmt.Println(numero)

	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalComPonteiro(&novoNumero) //& --> desse modo jogue o endereco de memoria onde esta variavel esta salva
	fmt.Println(novoNumero)
}
