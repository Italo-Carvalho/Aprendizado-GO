package main

import "fmt"

func main() {
	funcPorletra := map[string]map[string]float64{ // Map aninhado
		"G": {
			"Grabiela Silsa": 1121.20,
			"Guga Pereira":   8451.45,
		},
		"J": {
			"José João": 41502.20,
		},
		"I": {
			"Italo Carvalho": 99999.99,
			"Iuri Costa":     1548.80,
		},
	}

	delete(funcPorletra, "J")

	for letra, funcs := range funcPorletra {
		fmt.Printf("%s)\n", letra)
		for nome, salario := range funcs {
			fmt.Printf(" - %s (Salário: %v)\n", nome, salario)
		}
	}
}
