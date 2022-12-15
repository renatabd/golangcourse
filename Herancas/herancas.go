package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa    //coloca o nome do outro struct e nao passa um tipo
	curso     string
	faculdade string
}

//Em Go nao é possivel criar uma variavel e nao usar, mas struct ele permite.
func main() {
	fmt.Println("Herança")

	p1 := pessoa{"Renata", "Diniz", 33, 170}
	fmt.Println(p1)

	e1 := estudante{p1, "Nutrição", "UECE"}
	fmt.Println(e1)
	fmt.Println(e1.nome) //Ou e1.pessoa.nome - mas é redundante
}
