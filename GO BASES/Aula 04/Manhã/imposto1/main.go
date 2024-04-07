package main

import "fmt"

type preventError struct{
	//struct vazia
}

func (p preventError) Error() string{
	return "Erro: O salário digitado não está dentro do valor mínimo, tente novamente."
}

func main() {
	salario := 10000

	if salario < 15000{
		err := preventError{}
		fmt.Println(err.Error())
	} else{
		fmt.Println("Necessário pagamento de imposto")
	}

}