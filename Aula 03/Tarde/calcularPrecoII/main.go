package main

import (
	"fmt"
	"sync"
)

type produtos struct{
	nome 		string
	preco		float64
	quantidade	int
}

type services struct{
	nome			string
	preco			float64
	minTrabalhados	float64
}

type manutencao struct{
	nome 		string
	preco		float64
}

func somarProdutos(produtos []produtos)float64{
	precoTotal := 0.0
	for _, p := range produtos{
		precoTotal += float64(p.preco * float64(p.quantidade))
	}
	return precoTotal
}

func somarServices(services []services)float64{
	precoTotal := 0.0
	for _, s := range services{
		precoTotal += float64(s.preco * s.minTrabalhados) / 30
	}
	return precoTotal
}

func somarManutencao(manutencao []manutencao)float64{
	precoTotal := 0.0
	for _, m := range manutencao{
		precoTotal += m.preco
	}
	return precoTotal
}

func main(){
	produtos := []produtos{
        {nome: "Produto1", preco: 10.0, quantidade: 10},
    }
    servicos := []services{
        {nome: "Servico1", preco: 1.0, minTrabalhados: 180},
    }
    manutencoes := []manutencao{
        {nome: "Manutencao1", preco: 50.0},
    }

	var wait sync.WaitGroup
	wait.Add(3)

	results := make(chan float64, 3)

	go func(){
		defer wait.Done()
		results <- somarProdutos(produtos)
	}()

	go func(){
		defer wait.Done()
		results <- somarServices(servicos)
	}()

	go func(){
		defer wait.Done()
		results <- somarManutencao(manutencoes)
	}()

	go func(){
		wait.Wait()
		close(results)
	}()

	var totalTudo float64
	for valores := range results {
		totalTudo += valores
	}


	fmt.Printf("o valor total foi de: %.2f \n", totalTudo)
}