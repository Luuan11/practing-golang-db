package ordenar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdering(t *testing.T){
	numbers  := []int{1,3,4,5,6,7,8,9,2}
	expect   := []int{1,2,3,4,5,6,7,8,9}
	res := ordenar(numbers)
	assert.Equal(t, res, expect, "devem ser iguais")
}