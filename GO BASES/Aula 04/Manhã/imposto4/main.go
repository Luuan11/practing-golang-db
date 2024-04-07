package main

import (
	"errors"
	"fmt"
)

type empresa struct{
	funcionario 	    string
	horasTrabalhadas	int
	valorHora			int
}

func (e empresa) salarioMensal()(float64, error){
	salario := float64(e.horasTrabalhadas) * float64(e.valorHora)
	if salario >= 20000 {
		salarioImposto := salario * 0.10
		salario -= salarioImposto
		fmt.Printf("Seu salario foi de %.2f e terá que pagar imposto \n", salario)
	} else{
		fmt.Printf("Seu salario foi de %.2f e não precisar pagar imposto \n", salario)
	}
	
	if e.horasTrabalhadas < 80 && e.horasTrabalhadas > 0 {
		return 0, errors.New("erro: o trabalhador não pode ter trabalhado menos de 80 horas por mês")
	}

	return salario, nil
}

func main() {
	funcionario1 := empresa{
		funcionario: 	"Lucas",
		horasTrabalhadas: 90,
		valorHora: 200.0,
	}

	salario, err := funcionario1.salarioMensal()
	if err != nil{
		fmt.Println("Erro", err)
	} else {
		fmt.Printf("O salario do funcionário %s é de: %.2f Mensais \n", funcionario1.funcionario, salario)
	}

}