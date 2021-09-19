package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s %s (iteração  %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// fale("Maria", "pq vc não fala cmg?", 3)
	// fale("Jão", "Só posso falar depois de vc!", 1)

	/*
	 A go routine só funciona se a [thread] principal estive funcionando
	*/
	// go fale("Maria", "Ei...", 500)
	// go fale("João", "Opa...", 500)

	go fale("Maria", "Entendi!!!", 10)
	fale("Jão", "Parabéns!", 5)

}
