package main

import "fmt"

type Item struct {
	Nome       string
	Preco      int
	Quantidade float64
}

type Carrinho struct {
	Items []Item
}

func (c *Carrinho) DividirConta(emails []string) map[string]int {
	valorTotalDosItens := 0

	result := make(map[string]int)
	if len(emails) == 0 {
		return result
	}

	for _, item := range c.Items {
		valorTotalDosItens += item.Preco * int(item.Quantidade)
	}

	total := valorTotalDosItens / len(emails)
	resto := valorTotalDosItens % len(emails)

	for i, email := range emails {
		if i == 0 {
			result[email] = total + resto
		} else {
			result[email] = total
		}
	}
	return result
}

func main() {
	carrinho := Carrinho{
		Items: []Item{
			{"Banana", 75, 3},
			{"Frango", 150, 2},
			{"Caixa de chocolate", 1000, 2},
		},
	}

	emails := []string{"douglasfrango@gmail.com", "augustugalego@gmail.com", "deyvin.marketing@gmail.com"}

	fmt.Println(carrinho.DividirConta(emails))
}
