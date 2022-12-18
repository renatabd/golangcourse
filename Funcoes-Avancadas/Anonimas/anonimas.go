package main

import "fmt"

func main() {
	func(texto string) {
		fmt.Println(texto)
	}("Oi")

	//Outra forma:
	// func main() {
	// 	func() {
	// 		fmt.Println("Oi")
	// 	}()

	//Outra forma:
	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto) //Concatena informacoes
	}("Oi")
	fmt.Println(retorno)
}
