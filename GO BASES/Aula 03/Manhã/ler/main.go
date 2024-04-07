package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type produto struct{
	id 		int
	qtd 	int
	preco 	float64
}

func produtoLista(data []string) ([]produto,error) {
	produtos := make([]produto, 0, len(data))
	for i, row := range data{
		if i == 0{
			continue
		}

		var p produto
		fmt.Sscanf(row, "%d,%f,%d", &p.id, &p.preco, &p.qtd)
		produtos = append(produtos, p)
	}
	return produtos, nil
}

func lerArquivo(filename string)([]produto, error){
	arq, err := os.Open(filename)
	if err != nil{
		return nil, errors.New("Erro ao tentar abrir o arquivo informado")
	}

	defer arq.Close()

	scanner := bufio.NewScanner(arq)
	var data []string
	for scanner.Scan(){
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, errors.New("erro ao tentar ler o arquivo informado")
	}

	return produtoLista(data)
}

func main() {
	produtos, err := lerArquivo("../guardar/produtos.csv")
	if err != nil {
		fmt.Println("Erro")
		return
	}

	fmt.Printf("ID\tPreco\tQuantidade\n")
	var total float64 = 0.0
	for _, p := range produtos {
		fmt.Printf("%d\t%.2f\t%d\n", p.id, p.preco, p.qtd)
		total += p.preco * float64(p.qtd)
	}
	fmt.Printf("\t%.2f\n", total)
}
