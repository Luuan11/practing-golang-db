package main

import (
	"fmt"
	"os"
)

type produto struct{
	id 		int
	qtd 	int
	preco 	float64
}

func guardarArquivo(produto []produto, fileName string) {
	var valores []byte
	valores = append(valores, "ID,PRECO,QUANTIDADE\n"...)

	for _, produtoItem := range produto{
		valores = append(valores, fmt.Sprintf(" %d , %2.f , %d\n" , produtoItem.id, produtoItem.preco, produtoItem.qtd)...)
	}

	err := os.WriteFile(fileName+".csv", valores, 0644)
	if err != nil{
		fmt.Println("Erro")
		return
	}
}

func main() {
	produtos := []produto{
		{
			id:         1,
			qtd: 8,
			preco:      6.99,
		},
		{
			id:         2,
			qtd: 		20,
			preco:      12.99,
		},
		{
			id:         3,
			qtd: 		10,
			preco:      9.99,
		},
	}

	guardarArquivo(produtos, "produtos")
}
