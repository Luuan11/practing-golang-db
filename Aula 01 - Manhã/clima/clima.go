package main

import "fmt"

func main(){
	var (
		temperatura float64 = 26.3
		umidade float64 = 47.6
		pressaoAtmosferica float64 = 1016.23
	)

	fmt.Println("Você buscou por dados climaticos em SP, Segue os dados abaixo:")
	fmt.Printf("Temperatura: %.2f°C\n", temperatura)
    fmt.Printf("Umidade: %.2f%%\n", umidade)
    fmt.Printf("Pressão Atmosférica: %.2f hPa\n", pressaoAtmosferica)
}