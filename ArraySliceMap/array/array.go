package main

import "fmt"

/*
Ao usar o array em uma função é realizada
uma cópia para outro endereço de mémoria
*/
func main() {
	// homogênea (mesmo tipo) e estática (fixo)
	var notas [3]float64
	fmt.Println(notas)
	notas[0], notas[1], notas[2] = 7.8, 5.5, 9.4
	// notas[3] = 5.5 [INDEX ERROR]
	fmt.Println(notas)

	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}
	media := total / float64(len(notas))
	fmt.Printf("Média %.2f\n", media)

}
