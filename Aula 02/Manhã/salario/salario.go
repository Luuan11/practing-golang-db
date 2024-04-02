package main

import (
	"errors"
	"fmt"
)

func main() {
	categoria := "A"
	horas := 180.00

	if resultado, err := calcularSalario(categoria, horas) ; err != nil {
		panic(err.Error())
	} else {
		fmt.Printf("Sua categoria é %s por hora, sendo assim: salario de %.2f\n", categoria, resultado)
	}
}

func calcularSalario(categoria string, horas float64) (float64, error) {

	juros := 1.0
	salario := 0

	if categoria == "C" {
		salario = 1000
	} else if categoria == "B" {
		salario = 1500
		if horas > 160 {
			juros = 1.2
		}
	} else if categoria == "A" {
		salario = 3000
		if horas > 160 {
			juros = 1.5
		}
	} else {
		return 0.0, errors.New("Erro, categoria invalida")

		
	}
	return float64(horas) * float64(salario) * juros, nil

}
