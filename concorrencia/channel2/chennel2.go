package main

import (
	"fmt"
	"time"
)

// Channel (canal) - é a forma de comunicação entre goroutines
// channel é um tipo
// O channel força um sincronismo e só continua o programa quando um dado é recebido

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base // enviando dados para  o channel
	time.Sleep(time.Second)
	c <- 3 * base
	time.Sleep(time.Second * 3)
	c <- 4 * base
}

func main() {
	c := make(chan int)
	go doisTresQuatroVezes(2, c)
	a, b := <-c, <-c // recebendo dados do canal
	fmt.Println(a, b)
	fmt.Println(<-c)
	// fmt.Println(<-c) [ERROR](deadlock) esse channel não tem nenhum valor a receber
}
