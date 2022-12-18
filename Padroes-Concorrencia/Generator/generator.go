//GENERATOR: uma função que vai encapsular uma chamada pra uma GoRoutine e devolver um canal.
//Usa-se para diminuir a complexidade.
package main

import (
	"fmt"
	"time"
)

func main() {

	canal := escrever("Olá Mundo!")

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}

}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()
	return canal
}
