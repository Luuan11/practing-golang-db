package main

import (
	"fmt"
	"errors"
)

// type preventError struct{}

func exibirErro() error{
	return errors.New("Erro: O salário digitado não está dentro do valor mínimo, tente novamente.")
} 

func main() {
	salario := 10000

	if salario < 15000{
		err := exibirErro()
		fmt.Println(err.Error())
	} else{
		fmt.Println("Necessário pagamento de imposto")
	}

}