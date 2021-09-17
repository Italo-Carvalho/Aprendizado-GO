package main

import "fmt"

type item struct {
	produtoID int
	qtde      int
	preco     float64
}

type pedido struct {
	userID int
	itens  []item // slice de item
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.qtde)
	}
	return total
}

func main() {
	pedido := pedido{
		userID: 1,
		itens: []item{
			{1, 2, 12.10}, // item{id, qtd, preco}
			{2, 1, 23.48},
			{11, 100, 1.99},
		},
	}

	fmt.Printf("Valor total do pedido é R$ %.2f", pedido.valorTotal())
}
