package main

import (
	"fmt"
	"errors"
)

const (
	dog 		= "dog"
	cat 		= "cat"
	hamster 	= "hamster"
	tarantula  = "tarantula"
 )

func animal(texto string) (func(int) float64, error){
	var alimento float64

	switch texto {
	case dog:
		alimento = 10
	case cat:
		alimento = 5
	case hamster:
		alimento = 0.250
	case tarantula:
		alimento = 0.150
	default:
		return nil, errors.New("Erro, animal inválido, verifique e tente novamente")
	}

	return func(contador int) float64 {
		return float64(contador) * alimento
	}, nil
}


func main(){
	animalDog, err := animal(dog)
	animalCat, err := animal(cat)
	animalHamster, err := animal(hamster)
	animalTarantula, err := animal(tarantula)
	if err != nil {
		panic(err.Error())
	}
	amount := 0.0
	amount += animalDog(5)
	amount += animalCat(8)
	amount += animalHamster(12)
	amount += animalTarantula(13)

	fmt.Printf("Vamos precisar de %2.f kilos de ração para todos os bichinhos", amount)
}

