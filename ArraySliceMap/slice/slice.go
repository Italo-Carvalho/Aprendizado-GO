package main

import (
	"fmt"
	"reflect"
)

func main() {
	// array [tamanho fixo], slice [tamanho variavel]
	a1 := [3]int{1, 2, 3} // array
	s1 := []int{1, 2, 3}  //slice
	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}
	// Slice não é um array!, Slice define um pedaço de um array.
	s2 := a2[1:3] //[x:y], x e um ponteiro para uma posição do array, y serve pra determinar o tamanho do slice
	fmt.Println(a2, s2)

	s3 := a2[:2]
	fmt.Println(a2, s3)

	// Slide de um Slice
	s4 := s2[:1]
	fmt.Println(s2, s4)
}
