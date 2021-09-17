package main

import "fmt"

func inc1(n int) {
	n++
}

// um ponteiro é representado por um [*]
func inc2(n *int) {
	// [*] é usado para desreferenciar, ou seja,
	// ter acesso ao valor no qual o ponteiro aponta
	*n++
}
func main() {
	n := 1
	inc1(n) // é criado uma copia e não altera o valor da variavel passada
	fmt.Println(n)
	// [&] é usado para obter o endereço da variavel
	inc2(&n) // por está passando o endereço a função altera a variavel passada
	fmt.Println(n)
}
