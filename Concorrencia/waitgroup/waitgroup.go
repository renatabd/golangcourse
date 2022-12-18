package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitgroup sync.WaitGroup

	waitgroup.Add(2) //Duas goroutines que fazem parte do grupo de espera/fila

	go func() {
		escrever("Olá mundo")
		waitgroup.Done() //-1
	}()
	go func() {
		escrever("Programando em GO")
		waitgroup.Done() //-1
	}()

	waitgroup.Wait() //esperar a contagem das goroutines chegar em zero. Ou seja espera o Done do primeiro grupo, depois o segundo.
	//Vão ser executadas de forma concorrente.

}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
