package main

import "fmt"

//Defer: adia a execução de um determinado pedaço de código

func funcao1() {
	fmt.Println("Executando a funcao 1")
}

func funcao2() {
	fmt.Println("Executando a funcao2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada. Resultada será retornado!")
	fmt.Println("Entrando na funcao para verificar se o aluno esta aprovado")
	media := (n1 + n2) / 2
	if media >= 6 {
		// fmt.Println("Média calculada. Resultada será retornado") <-- eu poderia fazer repetindo o print aqui e depois
		return true
	}
	// fmt.Println("Média calculada. Resultada será retornado")
	return false
}

func main() {
	// defer funcao1() //DEFER = ADIAR, ele adia a funcao 1 para depois das outras funcoes, sera a ultima a ser executada
	// funcao2()
	// fmt.Println("oi")

	fmt.Println(alunoEstaAprovado(7, 8))
}

//Útil quando estamos lidando um banco de dados! Quando se chama uma funcao abre-se uma conexao com o banco de dados. Pra fechar uma funcao depois de cada insert pode-se utilizar o defer no lugar de ficar repetindo o apos o return.
