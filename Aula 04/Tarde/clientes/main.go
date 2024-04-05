package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

type file struct{
	filename string
}

func (f file)lerArquivo()(string, error){
	ler, err := os.Open(f.filename)
	if err != nil {
		return "", errors.New("O arquivo indicado não foi encontrado ou está danificado")
	}
	defer ler.Close()

	conteudo, err := io.ReadAll(ler)
	if err != nil{
		return "", err
	}
	return string(conteudo), nil
	
}

func main() {
	arquivo := file{filename: "customers.txt"}

	conteudo, err := arquivo.lerArquivo()
	if err != nil{
		panic(err.Error())
	}

	fmt.Println(conteudo)
}