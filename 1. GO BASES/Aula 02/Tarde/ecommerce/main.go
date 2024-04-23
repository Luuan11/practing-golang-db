package main

import(
	"fmt"
	"errors"
)

const (
	Pequeno = "Pequeno"
	Medio= "Medio"
	Grande = "Grande"
)

type produto struct{
	tipoProduto string
	nomeProduto string
	precoProduto float64
}

type Produto interface{
	CalcularCusto() float64
}

type Ecommerce interface{
	total() (float64, error)
	Adicionar(p produto) 
}

func novoProduto(nome string, preco float64, tipoProduto string) produto{
		return produto{
			nomeProduto: nome,
			precoProduto: preco,
			tipoProduto: tipoProduto,
		}
}

type loja struct{
	produtosGerais []produto
}

func (l loja)total() (float64, error){
	res := 0.0
	for _, pValue := range l.produtosGerais{
		custo, err := pValue.CalcularCusto()
		if err != nil {
			return 0.0, errors.New("Erro")
		}
		res += custo
	}
	return res, nil
}

func (l *loja) Adicionar(p produto) {
	l.produtosGerais = append(l.produtosGerais, p)
}

func NovaLoja() Ecommerce {
	return &loja{}
}

func (p produto) CalcularCusto() (float64, error){
	switch p.tipoProduto{
	case Pequeno:
		return p.precoProduto * 1.00,nil
 	case Medio:
		return p.precoProduto * 1.03,nil
	case Grande:
		return p.precoProduto * 1.6 + 2500.00 ,nil
	default:
		return 0 , errors.New("valor incorreto")
	}
}
func main(){
	produto1 := novoProduto("Produto 1", 100, Pequeno)
	produto2 := novoProduto("Produto 2", 200, Medio)
	produto3 := novoProduto("Produto 3", 300, Grande)

	l := &loja{}
	l.Adicionar(produto1)
	l.Adicionar(produto2)
	l.Adicionar(produto3)

	total, err := l.total()
	if err != nil {
		fmt.Println("Erro ao calcular o total da loja:", err)
		return
	}

	fmt.Printf("Total da loja: $%.2f\n", total)
}