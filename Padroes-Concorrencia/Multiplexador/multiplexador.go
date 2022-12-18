//MULTIPLEXADOR: pega mais de dois canais e junta em um canal só
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	canal := multiplexar(escrever("Olá Mundo"), escrever("Programando em Go!"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}

}

func multiplexar(canalDeEntrada1, canalDeEntrada2 <-chan string) <-chan string {
	canaldeSaida := make(chan string)
	go func() {
		for {
			select {
			case mensagem := <-canalDeEntrada1:
				canaldeSaida <- mensagem
			case mensagem := <-canalDeEntrada2:
				canaldeSaida <- mensagem
			}
		}
	}()

	return canaldeSaida
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000))) //Número pseudoaleatório. Retirado 500 e substituido por time.Duration(rand.Intn(2000)
		}
	}()
	return canal
}
