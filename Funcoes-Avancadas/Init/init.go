package main

import "fmt"

var n int

func init() { //Funcao que vai ser executada antes da funcao main
	//Pode ter uma funcao init por arquivo e nao por pacote
	fmt.Println("Executando a funcao init")
	n = 10

}

func main() {
	fmt.Println("Funcao main sendo executada")
	fmt.Println(n)
}
