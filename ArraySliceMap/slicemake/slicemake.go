package main

import "fmt"

func main() {
	s := make([]int, 10) // slice de inteiro com 10 elementos
	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 20)        // slice com 10 elementos, com um array interno de 20 posições
	fmt.Println(s, len(s), cap(s)) // cap(), pegar capacidade máxima desse slice

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s)) // 20 elementos no slice, array interno 20

	/*
		21 elementos no slice, e o array dobrou de tamanho
		para 40 por ter atingido o seu limite antigo de 20 posições
	*/
	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))
}
