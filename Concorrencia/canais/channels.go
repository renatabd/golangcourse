package main

import (
	"fmt"
	"time"
)

func main() {
	//Canais: canais de comunicação. Envia e recebe dados para sincronizar as goroutines
	canal := make(chan string) //chan = canal, significa que eu só posso enviar e receber dados do tipo string
	go escrever("Olá Renata", canal)

	fmt.Println("Depois da função escrever começar a ser executada")

	for { //infinito, mas nossa funcao só repete 5 vezes, então da DEADLOCK

		mensagem, aberto := <-canal //canal vai *esperar* receber um valor (primeiro setinha, depois canal)
		if !aberto {                //avalia se o canal ainda está aberto, e ai sai desse loop
			break //comando break - sai do loop infinito
		}
		fmt.Println(mensagem)
	}

	for mensagem := range canal { //Para cada mensagem recebida no canal, vai printar na tela mensagem
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa")
}

//NATUREZA DOS CANAIS: um canal tem duas operações: enviar e receber dados.
//Essas operações bloqueam a execução do programa

func escrever(texto string, canal chan string) {
	// time.Sleep(time.Second * 5) //espera por 5 segundos
	for i := 0; i < 5; i++ { //o canal nao esperou as 5 execuções, a partir do momento que o canal recebe a mensagem, ele termina o programa
		// fmt.Println(texto) - anterior, novo:
		canal <- texto //enviando um valor: primeiro canal depois setinha
		time.Sleep(time.Second)
	}

	close(canal) //após a 5x, ele fecha
}

// ❯ go run channels.go
// Depois da função escrever começar a ser executada
// Olá Renata
// fatal error: all goroutines are asleep - deadlock! --> cuidado

// goroutine 1 [chan receive]:
// main.main()
//         /Users/rbelizario/Documents/GOLANGCOURSE/Concorrencia/canais/channels.go:17 +0xc8
// exit status 2
//DEADLOCK: o programa trava pq ta esperando algo que nunca irá acontecer.
