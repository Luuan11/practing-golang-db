package main

import "fmt"

// type preventError struct{}

func exibirErro(salario int) error{
	return fmt.Errorf("erro: o mínimo tributável é 15.000 e o salário informado é: %d", salario)
} 

func main() {
	salario := 10000

	if salario < 15000{
		err := exibirErro(salario)
		fmt.Println(err.Error())
	} else{
		fmt.Println("Necessário pagamento de imposto")
	}

}
