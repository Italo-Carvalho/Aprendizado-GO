package main

import "fmt"

func imprimirAprovados(aprovados ...string) {
	fmt.Println("Lista de Aprovados")
	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i, aprovado)
	}
}

func main() {
	aprovados := []string{"Ítalo", "João", "Maria", "Valentina"}
	imprimirAprovados(aprovados...)
}
