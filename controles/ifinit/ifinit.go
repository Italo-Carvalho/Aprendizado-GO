package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano()) // gerar um numero aleatorio com base no nano segundo
	r := rand.New(s)                           // gerar um numero novo toda vez
	return r.Intn(10)                          // limitar o rand numero ate 10

}

func main() {
	if i := numeroAleatorio(); i > 5 { // atribuição = [i:= numeroAleatorio()] [;] condição = [i > 5]  | tbm é suportado no switch
		fmt.Println("Ganhou")
		fmt.Println(i)
	} else {
		fmt.Println("Perdeu")
		fmt.Println(i)
	}
}
