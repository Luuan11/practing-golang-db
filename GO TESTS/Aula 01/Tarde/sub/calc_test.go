package calc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	num1 := 5
	num2 := 5
	expect := 10
	res := Sum(num1, num2)
	assert.Equal(t, res, expect, "devem ser iguais")
}

func TestSub(t *testing.T){
	num1 := 10
	num2 := 5
	expect := 5
	res := Sub(num1, num2)
	assert.Equal(t, res, expect, "devem ser iguais")
}

