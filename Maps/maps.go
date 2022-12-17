package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{ //Chaves (dentro do colchete) e valores todos do mesmo tipo, a nao ser que esteja usando o interfaces
		"nome":      "Renata",
		"sobrenome": "Diniz",
	}
	//Diferença pro struct: nao pode fazer usuario.nome --> (usuario["nome])
	usuario2 := map[string]map[string]string{ //aninhando maps
		"nome": {
			"primeiro": "Joao",
			"ultimo":   "Hennrique",
		},
		"curso": {
			"cursinho": "Publicidade",
			"campus":   "campus 1",
		},
	}
	fmt.Println(usuario)
	fmt.Println(usuario2)
	delete(usuario2, "nome") //deletar uma informação do println
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{ //adicionar uma informacao do println
		"nome": "Gemeos",
	}
	fmt.Println(usuario2)
}
