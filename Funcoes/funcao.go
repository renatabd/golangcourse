package main

import "fmt"

//Funçoes sao como receitas e bolo
//Parametros: sao dados voce coloca na funcao pra ela funcionar
//Retorno: o que a funcao devolve pra vc

//Dentro do parenteses: parametro
//Fora do parenteses: return
//Entre {}: a função a ser executada
func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

//Se vc nao declarar o tipo do parâmetro, ele vai pegar o do último. Aqui abaixo está um exemplo de uma func tendo dois retornos (é muito comum em Go e é MUITO usado!):
func calculosMatematicos(n1, n2 int16) (int16, int16) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func() {
		fmt.Println("Função f")
	}
	f()

	//Ou poderia ser assim:
	//var f = func(txt  string) {
	//	fmt.Println(txt)
	//}
	//f("Texto da função 1")

	//Ou poderia ser assim[2]:
	//var f = func(txt string) string {
	//	fmt.Println(txt)
	//return txt
	//}
	//resultado := f("Texnto da função 2")
	//fmt.println(resultado)

	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	//Se eu quiser ignorar uma variável, eu uso o símbolo do underline _ no lugar da varíavel
	//Exemplo: resultadoSoma, _ := calculosMatematicos(10, 15)
	//fmt.Println(resultadoSoma)
}
