package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.")
	fmt.Println(" Nova")
	fmt.Println("linha")

	x := 3.141516

	// fmt.Println("O valor de x é" + x)

	xs := fmt.Sprint(x) //retorna uma string

	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de x é", x)

	fmt.Printf("O valor de x é %.2f", x) //O .2 antes do f (de float) serve para imprimir apenas duas cadas decimais

	a := 1
	b := 1.999
	c := false
	d := "opa"
	fmt.Printf("\n %d %f %.1f %t %s", a, b, b, c, d) //%d [inteiro], %f [float], %t [boolean], %s [string]
	fmt.Printf("\n %v %v %v %v", a, b, c, d)         //%v [É generico e serve para varios tipos de variaveis]
}
