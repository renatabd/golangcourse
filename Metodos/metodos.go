//Metodos estao associados a alguma coisa
//METODOS = ACOES DAS CLASSES
package main

import "fmt"

// func escrever() {
// 	fmt.Println("Escrevendo...")
// }

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() { //u = variavel para referenciar outros campos do mesmo usuario que chamou esse metodo. Pode ser qualquer letra, palavra etc. A unica coisa que tem que ser igual é o tipo
	// fmt.Println("Dentro do método salvar")
	fmt.Printf("Salvando os dados do Usuário %s no banco de dados\n", u.nome) //%s string %d numero %f numero inteiro
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	// escrever()
	usuario1 := usuario{"Usuario 1", 20}
	usuario1.salvar() //Preciso criar um método para dar acoes! Por exemplo, salvar dados no banco de dados.

	usuario2 := usuario{"Renata", 33}
	usuario2.salvar()

	maiorDeIdade := usuario2.maiorDeIdade()
	fmt.Println(maiorDeIdade)

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
}
