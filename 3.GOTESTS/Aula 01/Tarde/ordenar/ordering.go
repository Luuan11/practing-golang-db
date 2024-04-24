package ordenar

import "slices"

func ordenar(numbers[]int) []int {
	Clonadosnumbers := slices.Clone(numbers)
	slices.Sort[[]int](Clonadosnumbers)
	return Clonadosnumbers
}