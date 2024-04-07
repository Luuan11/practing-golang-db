package main

import (
	"fmt"
	"errors"
	"slices"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func valores (valor ...int)(res int){
	for _, valor := range valor{
		res += valor
	}
	return
}
func operation(txt string) (func(valor ...int) int, error) {
	switch txt{
	case "minimum":
		return func(valor ...int) int{ return slices.Min(valor)}, nil
	case "average":
		return func(valor ...int) int{ return valores(valor...) / len(valor)} ,nil
	case "maximum":
		return func(valor ...int) int{ return slices.Max(valor)} ,nil
	default:
		return nil, errors.New("Valor incorreto, as opções que temos é, minimum, average e maximum, tente novamente")
	}
}

func main(){
	minhaFunc, err := operation(minimum)
	averageFunc, err := operation(average)
	maxFunc, err := operation(maximum)

	if err != nil {
		panic("erro na aplicação, tente novamente")
	} else {
		minValue := minhaFunc(2, 3, 3, 4, 10, 2, 4, 5)
		averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
		maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

		fmt.Println("minValue: ", minValue)
		fmt.Println("averageValue: ", averageValue)
		fmt.Println(" maxvalue: ", maxValue)
	}
}
