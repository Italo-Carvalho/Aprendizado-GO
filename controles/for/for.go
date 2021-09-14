package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	for i <= 10 { // While like
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 { // For like
		fmt.Printf("%d ", i)
	}

	fmt.Println("\nMisturando... ")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Print("Par ")
		} else {
			fmt.Print("Impar ")
		}
	}

	for {
		//laço infinito
		fmt.Println("Para sempre...")
		time.Sleep(time.Second * 5) // 5 segundos
	}
}
