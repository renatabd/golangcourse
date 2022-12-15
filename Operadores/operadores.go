package main

import "fmt"

func main() {
	// ARITMETICOS
	//+, -, /, *, % (MOD/remainder)
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 2
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 3

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	//EXEMPLO:
	//var n1 int16 = 10
	//var n2 int32 = 25
	//soma := n1 + n2
	//fmt.println(soma)
	//O exemplo acima não vai funcionar pois não se pode fazer nada no Go com variaveis que sao de tipos diferentes

	var n1 int16 = 10
	var n2 int16 = 25
	soma2 := n1 + n2
	fmt.Println(soma2)

	//FIM ARITMÉTICOS

	//ATRIBUIÇÃO -> =
	var variavel1 string = "string"
	variavel2 := "string2"
	fmt.Println(variavel1, variavel2)
	//FIM ATRIBUIÇÃO

	//OPERADORES RELACIONAIS (retornam apenas resultados de true/false)
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2) //!= significa diferente
	//FIM DOS RELACIONAIS

	//OPERADORES LÓGICOS (E, OU E NÃO)
	fmt.Println("-------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso) //Avalia todos os operadores e avalia se são verdadeiros (E)
	fmt.Println(verdadeiro || falso) //OU --> ||
	fmt.Println(!verdadeiro)         //Negação ! <-- exclamação
	fmt.Println(!falso)
	//FIM DOS OPERADORES LÓGICOS

	//OPERADORES UNÁRIOS - Só faz a interação com uma variável por vez
	numero := 10
	numero++ // numero + 1
	numero += 50
	numero /= 3
	fmt.Println(numero)
	//numero-- (subtrair)
	//numero -=20
	//numero *= 3
	//numero /= 10
	//numero %= 3 - restante da divisão
	//FIM DOS OPERADORES UNÁRIOS

	//OPERADOR TERNÁRIO - Não existe isso em GO
	//Colocar o valor de uma variável em geral em uma função
	//texto := numero > 5 ? "Maior que 5" : "Menor que 5"
	//fmt.Println(texto)
	//A premissa do GO é que só exista uma maneira de fazer cada coisa

	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}
	fmt.Println(texto)
}
