# Teste Técnico — Stone
 
Implementação de um carrinho de compras com divisão de conta entre pessoas.
 
## Como rodar
 
```bash
go run main.go
```
 
## O que foi implementado
 
### `DividirConta(emails []string) map[string]int`
 
Método do `Carrinho` que divide o valor total dos itens igualmente entre os emails fornecidos.
 
- Calcula o valor total multiplicando o preço pela quantidade de cada item
- Divide o total pelo número de pessoas
- O resto da divisão é adicionado ao primeiro email da lista
- Retorna um `map[string]int` com cada email e o valor que deve pagar
- Retorna um map vazio se nenhum email for fornecido
## Exemplo
 
```go
carrinho := Carrinho{
    Items: []Item{
        {"Banana", 75, 3},
        {"Frango", 150, 2},
        {"Caixa de chocolate", 1000, 2},
    },
}
 
emails := []string{
    "douglasfrango@gmail.com",
    "augustugalego@gmail.com",
    "deyvin.marketing@gmail.com",
}
 
fmt.Println(carrinho.DividirConta(emails))
// map[augustugalego@gmail.com:825 deyvin.marketing@gmail.com:825 douglasfrango@gmail.com:825]
```
