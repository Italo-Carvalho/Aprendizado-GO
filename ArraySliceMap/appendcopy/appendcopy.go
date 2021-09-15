package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int

	// array1 = append(array1, 1, 4, 5, 6) [ERROR] o primeiro argunmento tem que ser um slice
	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2)
	copy(slice2, slice1) // copiar tudo do slice1 para o slice2
	/*
		[! Aternção, o copy não faz que o slice cresça]
		por contar do slice2 ter apenas 2 de espaço, foi copiado
		apenas os 2 elementos do inicio do slice1 para o slice2
	*/
	fmt.Println(slice2)

}
