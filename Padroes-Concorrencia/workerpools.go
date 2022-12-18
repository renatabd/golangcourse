//WORKER POOLS: Uma fila de tarefa para serem executadas - e você tem varios workers (processos) pegando itens dessa fila de maneira independente.
package main

import "fmt"

func main() {
	tarefas := make(chan int, 45)    //Ter a sequência de numeros que vai ser preciso calcular
	resultados := make(chan int, 45) //Armazena os resultados a medida que eles forem sendo calculados

	go worker(tarefas, resultados) //POSSO REPETIR A GO ROUTINE PRA IR MAIS RÁPIDO, MAS EM UM DADO MOMENTO VAI FICAR LIMITADA A POTENCIA DA PRÓPRIA MÁQUINA

	for i := 0; i < 45; i++ {
		tarefas <- i
	}
	close(tarefas)

	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}
}

//FUNÇÃO WORKER
func worker(tarefas <-chan int, resultados chan<- int) { //tarefas só recebe <-chan e resultados só envia chan<-
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
