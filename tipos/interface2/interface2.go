package main

import "fmt"

/*
	É mais comum, que no nivel de interface vc não gere alterações
	no obj no qual vc está trabalhando, e sim em ler a coisas de  uma
	forma que não gere nenhum efeito colateral no obj
*/
type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	model           string
	turboLigado     bool
	velocidadeAtual int
}

func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
}

func main() {
	// nivel de estrutura
	carro1 := ferrari{"F40", false, 0}
	carro1.ligarTurbo()

	// nivel de interface
	var carro2 esportivo = &ferrari{"F40", false, 0}
	carro2.ligarTurbo()

	fmt.Println(carro1, carro2)
}
