package main

import (
	"fmt"
	"math"
	"reflect"
)

/*
int64 = 8 bytes
int32 = 4 bytes
int = 8/4 bytes dependendo da sua arquitetura
array = um conjunto de dados do mesmo tipo com um tamanho fixo
*/
func main() {
	//números inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	//sem sinal (só positivos)... uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	// com sinal... int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento da tabela Unicode (int32)
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	// números reais [float32, float64]
	x := float32(49.99)
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99)) // float64

	// Boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// String
	s1 := "Olá mundin"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanha da string é", len(s1))

	// String com multiplas linhas
	s2 := `Fala 
    zezé
    como
    ce
    ta?
    `
	fmt.Println("O tamanho da string é ", len(s2))

	// char??
	char := 'a'
	fmt.Println("O tipo de char é", reflect.TypeOf(char))
	fmt.Println(char) // 97
}
