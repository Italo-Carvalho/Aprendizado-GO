package main

import "fmt"

func main() {
	// var aprovados map[int]string // [x]y | x = chave, y = valor
	//mapa devem ser inicializados
	aprovados := make(map[int]string)
	aprovados[12345678] = "Maria"
	aprovados[31545343] = "Pedro"
	aprovados[62186284] = "Ítalo"
	fmt.Println(aprovados)
	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF:%d)\n", nome, cpf)
	}

	fmt.Println(aprovados[62186284]) //[CHAVE]
	delete(aprovados, 31545343)      // delete(map, chave)
	fmt.Println(aprovados)
}
