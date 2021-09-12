package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4 //float64
	y := 2
	/* para calcular números de diferentes tipos
	   e preciso realizer a conversão para ingualar os tipos */
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// cuidado...
	fmt.Println("Teste " + string(97))

	// int para string
	fmt.Println("Test " + strconv.Itoa(123))
	// ou
	fmt.Println("Test " + fmt.Sprint(123))

	// string para int
	num, _ := strconv.Atoi("123") // strconv.Atoi("123") retornara dois valores
	// o [_] serve para ignora a variavel

	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("true") //apenas true e 1 retorna verdadeiro

	if b {
		println(b)
	}

}
