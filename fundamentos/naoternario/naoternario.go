package main

import "fmt"

// Não tem operador ternário
func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovador"

}
func main() {
	fmt.Println(obterResultado(6.2))
}
