package main

import "fmt"

func calcularImposto(salary float64) float64{
	if salary <= 50000 {
		return 0.0
	} else if salary <= 150000.00 {
		return salary * 0.17
	} 
	return salary * 0.27
}

func main(){
	fmt.Printf("imposto de até R$50.000: %.2f", calcularImposto(50000))
	// fmt.Printf("imposto de até R$150.000: %.2f", calcularImposto(150000))
	// fmt.Printf("imposto mais que R$150.000: %.2f", calcularImposto(150001))
}
