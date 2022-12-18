package main

import "fmt"

func recuperarExecucao() {
	// fmt.Println("Tentativa de recuperar a execução")
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso")
	}
}

func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A MÉDIA É EXATAMENTE 6") //Tem que passar uma mensagem dentro de panic
} //O panic mata a execução do programa

//NO TERMINAL:
// panic: A MÉDIA É EXATAMENTE 6

// goroutine 1 [running]:
// main.alunoEstaAprovado(...)
//         /Users/rbelizario/Documents/GOLANGCOURSE/funcoes-avancadas/panic-e-recover/panic-e-recover.go:14
// main.main()
//         /Users/rbelizario/Documents/GOLANGCOURSE/funcoes-avancadas/panic-e-recover/panic-e-recover.go:18 +0x30
// exit status 2

func main() {
	fmt.Println(alunoEstaAprovado(6, 6))
	fmt.Println("Pós execução")
}
