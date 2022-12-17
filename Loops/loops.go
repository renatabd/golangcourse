package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	//PRIMEIRO TIPO DE FOR
	for i < 10 { //for significa: enquanto isso for verdadeiro, algo ocorre. for = while
		i++ //De um em um
		fmt.Println("Incrementando i")
		time.Sleep(time.Second) //Dorme por 1 segundo --> dá uma pausa

	}
	fmt.Println(" ")

	//FOR PARECIDO COM IF INIT
	for j := 0; j < 10; j += 2 { //De 2 em 2
		time.Sleep(time.Second)
		fmt.Println("Adicionando j", j) //Pode-se passar variaveis dentro do Println
	}

	//FOR USANDO A CLAUSULA RANGE - iteração/repetição
	nomes := [3]string{"Renata", "Vitor", "Joao"} //Array
	for indice, nome := range nomes {             //Em for o primeiro parametro é sempre o indice e o segundo é sempre o valor!
		fmt.Println(indice, nome)
	}

	//ITERAÇÃO/LOOP/REPETIÇÃO POR STRING
	for indice, letra := range "RENATA" { //O iterável é a propria letra dentro da tabela ASCII
		fmt.Println(indice, string(letra))
	}

	//ITERAR POR MAP
	usuario := map[string]string{
		"nome":      "Renata",
		"sobrenome": "Diniz",
	}
	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
	//loop infinito
	for {
		fmt.Println("Oi")
		time.Sleep(time.Second)
	}
}
