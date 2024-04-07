package main

import "fmt"

func main() {
	meses := 2
	fmt.Println("Você digitou o número", meses)
	mes := mes(meses)
	fmt.Println("Resultado foi:")
	fmt.Printf(mes)
}

func mes(num int) string{
	var mes string

	switch {
	case num == 1:
		mes = "1 de Janeiro"
	case num == 2:
		mes = "2 de fevereiro"
	case num == 3:
		mes = "3 de março"
	case num == 4:
		mes = "4 de abril"
	case num == 5:
		mes = "5 de maio"
	case num == 6:
		mes = "6 de julho"
	case num == 7:
		mes = "7 de junho"
	case num == 8:
		mes = "8 de agosto"
	case num == 9:
		mes = "9 de setembro"
	case num == 10:
		mes = "10 de outubro"
	case num == 11:
		mes = "11 de novembro"
	case num == 12:
		mes = "12 de dezembro"
	default:
		return "Número inválido"
	}

	return mes
}