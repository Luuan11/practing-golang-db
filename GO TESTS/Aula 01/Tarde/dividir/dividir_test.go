package dividir

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividir(t *testing.T){
	num1 := 10
	num2 := 0
	expect := 0

	res, err := Dividir(num1, num2)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), ErroNumZero)
	assert.Equal(t, expect, res, "devem ser iguais")

}