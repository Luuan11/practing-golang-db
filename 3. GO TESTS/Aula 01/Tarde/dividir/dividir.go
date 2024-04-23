package dividir

import (
	"errors"
)

var ErroNumZero = "o Divisor näo pode ser 0"
var ErroBothZero = "ambos dos números estão zerados, verifique novamente"

func Dividir(num1, num2 int) (int, error) {
	if num2 == 0 && num1 == 0 {
		return 0, errors.New(ErroBothZero)
	} else if num2 == 0 {
		return 0, errors.New(ErroNumZero)
	}

	return num1 / num2, nil
}


