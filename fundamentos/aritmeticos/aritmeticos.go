package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	println("Soma =>", a+b)
	println("Subtração =>", a-b)
	println("Divisão =>", a/b)
	println("Multiplicação =>", a*b)
	println("Módulo =>", a%b)

	// bitwise
	println("E =>", a&b)   // 11 & 10 = 10
	println("Ou =>", a|b)  // 11 | 10 = 11
	println("Xor =>", a^b) // 11 ^ 10 = 01

	c := 3.0
	d := 2.0

	// outras operações usando math:
	fmt.Println("Maior =>", math.Max(float64(a), float64(b)))
	fmt.Println("Menor =>", math.Min(c, d))
	fmt.Println("Exponenciação =>", math.Pow(c, d))
}
