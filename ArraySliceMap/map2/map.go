package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"Ítalo Carvalho":  999999.99,
		"José João":       11325.45,
		"Grabriela Silva": 15456.75,
		"Pedro Junior":    1200.0,
	}
	funcsESalarios["Rafael Filho"] = 1350.0
	delete(funcsESalarios, "Inexistent") // Ao tenta excluir uma chave inexistente não acontece nada
	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
