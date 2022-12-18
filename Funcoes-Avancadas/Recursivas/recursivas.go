package main

import "fmt"

//Não sao muito utilizadas

func fibonacci(posicao uint) uint { //IMPORTANTE: As funcoes recursivas tem uma condicao de parada: tem que especificar pro programa o momento que ele tem que parar
	//estouro de pilha: se nao determinar a parada, forma uma pilha gigantesca e o sistema nao vai conseguir comportar (stackoverflow)
	if posicao <= 1 {
		return posicao //condição de parada
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

//Recursivas: funcoes que chamam elas mesmas. Depende de uma outra execucao dela mesma
//Não é recomendada em numeros muito grandes
func main() {
	fmt.Println("Funções Recursivas")
	//funcao que retorna um numero que ta em uma determinada posicao da sequencia de fibonacci
	//sequencia de fibonacci: o proximo numero é sempre a soma dos dois numeros anteriores
	//1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584
	posicao := uint(10)

	for i := uint(0); i < posicao; i++ {
		fmt.Println(fibonacci(i))
	}

	fmt.Println(fibonacci(posicao))
}
