package main

import "fmt"

func main() {
	canal := make(chan string, 2) //aqui estou colocando que o canal tem uma capacidade 2

	canal <- "Olá!"               //Vai dar deadlock, por isso se usa canais em funcoes separadas ou criar um canal com buffer.
	canal <- "Programando em go!" //Não vai printar esse
	// canal <- "Terceiro valor" //Deadlock

	mensagem := <-canal
	mensagem2 := <-canal
	fmt.Println(mensagem, ",", mensagem2)
}
