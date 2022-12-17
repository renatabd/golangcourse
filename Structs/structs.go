package main

import "fmt"

//O struct é uma vários/coleção de campos
//Cada campo tem um nome e um tipo
//Exemplo: salvar um usuário em uma variável
//Sintaxe básica do struct: type usuario struct
//Tem mais de uma maneira de declarar variaveis em struct
//Não da pra fazer range em struct

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

//Struct dentro de struct
type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Renata"
	u.idade = 33
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua Alodia", 200}

	user2 := usuario{"Rezinha", 35, enderecoExemplo}
	fmt.Println(user2)

	//Se eu quiser passar só um valor dos dois campos acima (nome e idade) eu posso fazer da seguinte forma:
	user3 := usuario{nome: "Renata"}
	fmt.Println(user3)

	user4 := usuario{idade: 33}
	fmt.Println(user4)

}
