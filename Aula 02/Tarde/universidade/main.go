package main

import "fmt"

type students struct {
	nome         string
	sobrenome    string
	rg           int
	dataAdmissao string
}

func (a *students) detalhesAluno() {
	fmt.Printf("Nome: %s %s\n", a.nome, a.sobrenome)
	fmt.Printf("RG: %d\n", a.rg)
	fmt.Printf("Data da entrada: %s\n", a.dataAdmissao)
}

func main() {
	students1 := students{
		nome: "Luan",
		sobrenome: "Fernando",
		rg: 	123456789 ,
		dataAdmissao: "18/03/2024",
	}

	students2 := students{
		nome: "Giovana",
		sobrenome: "Santos",
		rg:		123456789,
		dataAdmissao: "18/03/2023",
	}

	students1.detalhesAluno()
	students2.detalhesAluno()

}
