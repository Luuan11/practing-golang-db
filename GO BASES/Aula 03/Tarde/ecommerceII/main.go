package main

import "fmt"

type user struct{
	Nome	string
	Sobrenome string
	Email	string
	Produtos []*produtos
}

type produtos struct{
	Nome 	string
	preco	float64
	quantidade	int
}

func novoProduto(nome string, preco float64, quantidade int) *produtos{
	return &produtos{
		Nome: nome,
		preco: preco,
		quantidade: quantidade,
	}
}

func (u *user)addProduto(produtos *produtos, quantidade int){
	prodExist := false

	for _, p := range u.Produtos{
		if p == produtos{
			p.quantidade += quantidade
			prodExist = true
			break
		}
	}
	if !prodExist{
		produtos.quantidade = quantidade
		u.Produtos = append(u.Produtos, produtos)
	}
}

func (u *user) deletaProdutos(){
	u.Produtos = nil
}

func main(){
	user := &user{
		Nome: "Luan",
		Sobrenome: "Fernando",
		Email: "luan.fs@gmail.com",
	}

	ptd1 := &produtos{
		Nome: "Mouse sem fio",
		preco: 100,
	}

	user.addProduto(ptd1, 23)

	for _, p := range user.Produtos{
		fmt.Printf("Nome: %s, Preco: R$%2.f, Quantidade: %d Unidades \n", p.Nome, p.preco, p.quantidade)
	}
}