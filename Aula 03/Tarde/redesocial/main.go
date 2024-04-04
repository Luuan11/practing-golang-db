package main

import "fmt"

type user struct{
	Nome	string
	Sobrenome	string
	Idade		int
	Email		string
	Senha		string	
}

func (u *user) mudarNome(Nome, Sobrenome string ){
	u.Nome = Nome
	u.Sobrenome = Sobrenome
}

func (u *user) mudarIdade(Idade int){
	u.Idade = Idade
}

func (u* user)mudarEmail(Email string){
	u.Email = Email
}

func (u* user)mudarSenha(Senha string){
	u.Senha = Senha
}

func novoUser() *user {
	return &user{}
}

func main(){
	user := novoUser()
	user.mudarNome("Luan", "Fernando")
	user.mudarIdade(21)
	user.mudarEmail("luan.fs@gmail.com")
	user.mudarSenha("BatataQuente")

	fmt.Println("Dados do usuário")
	fmt.Printf("Nome e sobrenome: %s %s\n", user.Nome,user.Sobrenome)
	fmt.Printf("Idade: %d anos \n", user.Idade)
	fmt.Printf("Email: %s \n", user.Email)
}