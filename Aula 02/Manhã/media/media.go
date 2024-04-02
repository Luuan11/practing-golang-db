package main

import "fmt"

func calcularMedia(notas ...float64) float64{

	media := 0.0
	for _,n := range notas{
		media += n
	}
	return media / float64(len(notas))
}

func main(){
	mediaGeral := calcularMedia(10,10,5,8)
	fmt.Println("sua média foi",mediaGeral)
	notaAprovacao := 6.5

	if mediaGeral >= notaAprovacao{
		fmt.Println("Voce foi aprovado!")
	} else {
		fmt.Println("Parece que sua nota foi menor do que 6, infelizmente está de recuperação")
	}
}