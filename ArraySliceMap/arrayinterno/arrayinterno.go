package main

import "fmt"

func main() {
	s1 := make([]int, 10, 20)
	s2 := append(s1, 1, 2, 3) // s2 = s1[0,0,0...] + 1, 2, 3
	fmt.Println(s1, s2)

	s1[0] = 7 // por apontar para o mesmo endereço de mémoria tanto o s1 e s2 mudou o primeiro elemento
	fmt.Println(s1, s2)
}
