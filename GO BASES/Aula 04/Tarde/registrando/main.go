package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

type cliente struct{
	nome	string
	sobrenome	string
	rg			string
	telefone 	string
	endereco	string
}

type file struct{
	filename string
}

func novoID()(string, error){
	id := rand.Intn(100)
	if id == 0{
		return "", errors.New("ID gerado não pode ser Zero")
	}
	return fmt.Sprintf("Arq%d", id), nil
}

func (f file)lerArquivo()(string, error){
	_, err := os.Open(f.filename)
	if err != nil{
		return "", errors.New("erro: o arquivo indicado não foi encontrado ou está danificado")
	}
	return "" ,nil
}

func (c cliente) novoClient() *cliente{
	return &cliente{
		nome: c.nome,
		sobrenome: c.sobrenome,
		rg: c.rg,
		telefone: c.telefone,
		endereco: c.endereco,
	}
}

func main() {
	id, err := novoID()
	if err != nil{
		panic(err)
	}

	arquivo := file{filename: "customers.txt"}
	_, err = arquivo.lerArquivo()
	if err != nil{
		fmt.Println("Erro", err)
		panic(err)
	}

	novoCliente1 := cliente{
		nome: "Luan",
		sobrenome: "Chaves",
		rg: "teste",
		telefone: "teste",
		endereco: "teste",
	} 

	fmt.Printf("Seja bem vindo %s %s \n agora você tem acesso ao sistema \n", novoCliente1.nome, novoCliente1.sobrenome)
	fmt.Printf("Seu id é %s \n", id)
}