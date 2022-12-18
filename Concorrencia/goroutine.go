package main

import (
	"fmt"
	"time"
)

func main() {
	//CONCORRENCIA != Paralelismo
	//Paralelismo: duas tarefas que estao sendo executadas exatamente ao mesmo tempo (quando se tem um processador com mais de um nucleo)
	//Concorrencia: não necessariamente estão sendo executadas ao mesmo tempo. Uma tarefa não espera a outra acabar para rodar também. Há um revezamento.
	// escrever("Olá Mundo") desse jeito, a linha de baixo nunca será executada
	go escrever("Olá mundo") // goroutine: um método que faz o revezamento (inicia mas nao espera ele terminar)
	escrever("Programando em GO")
}

func escrever(texto string) {
	for {
		fmt.Println("texto")
		time.Sleep(time.Second)
	}
}
