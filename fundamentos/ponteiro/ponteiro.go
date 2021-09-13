package main

import "fmt"

func main() {
	i := 1

	var p *int = nil // nil = null, o [*] serve para pegar o valor guardado neste entereço
	p = &i           // O [&] serve para coletar o endereço da variável
	*p++             // desreferenciando (pegando o valor)
	i++
	// Go não tem aritmética de ponteiros
	// p++ [ERROR]
	fmt.Println(p, *p, i, &i)
}
