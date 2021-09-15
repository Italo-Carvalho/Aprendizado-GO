package main

import "fmt"

func main() {
	numeros := [...]int{2, 4, 8, 11, 25} // compilador conta!
	for i, numero := range numeros {     // [i] retorna o indice, [numero] retorna o elemento
		fmt.Printf("%d) %d\n", i, numero)
	}
	for _, num := range numeros { // [_] ignora a variavel que retorna o indice
		fmt.Println(num)
	}

	for i := range numeros {
		fmt.Printf("%d) %d\n", i+1, numeros[i])
	}

}
