package main

import "fmt"

func generica(interf interface{}) { //Não colocando nada entre as chaves, tudo atenderia a interface
	fmt.Println(interf)
}

func main() {
	generica("String")
	generica(1)
	generica(true)

	mapa := map[interface{}]interface{}{ //UMA BAGUNÇA!
		1:            "string",
		float32(100): true,
		"String":     "String",
		true:         float64(12),
	}
	fmt.Println(mapa)
}
