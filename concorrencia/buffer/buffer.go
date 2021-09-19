package main

import (
	"fmt"
	"time"
)

/*
	Um channel sem buffer, quando um dado é enviado,
	aquela linha de código vai esperar até que aquele dado seja consumido.

	Um channel com buffer, você consegue enviar varios dados até que o buffer fique cheio,
	quando o buffer enche ai sim ele gera o bloqueio. quando o buffer esvazia ele tbm vai gerar
	um bloqueio esperando receber dados.
*/

func rotina(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	fmt.Println("Executou!")
	ch <- 6
}

func main() {
	ch := make(chan int, 3) // channel com um buffer de 3 posições
	go rotina(ch)
	time.Sleep(time.Second)
	fmt.Println(<-ch)
}
