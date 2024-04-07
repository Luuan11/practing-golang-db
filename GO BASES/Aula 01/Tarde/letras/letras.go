package main

import "fmt"

func main() {
	palavra := "Pneumoultramicroscopicossilicovulcanoconiótico"

	fmt.Printf("A palavra tem \"%s\" tem %d letras. \n" , palavra, len(palavra))

	soletrar(palavra)
}

func soletrar(palavra string) {
	fmt.Println("Soletrando a palavra")
	for _, letra := range palavra{
		fmt.Printf("%c-", letra)
	}
	fmt.Println()
} //esse foi hard