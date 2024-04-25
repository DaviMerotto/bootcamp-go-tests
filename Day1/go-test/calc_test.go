package gotest

import "testing"

func TestSum(t *testing.T) {
	if Sum(1, 1) != 2 {
		t.Errorf("a função Sum retornou um valor inesperado para a entrada 1,1")
	}
	if Sub(1, 1) != 0 {
		t.Errorf("a função Sub retornou um valor inesperado para a entrada 1,1")
	}

	result, err := Div(1, 0)
	if result != 1 || err != nil {
		if err != nil {
			t.Errorf("%v", err)
		} else {
			t.Errorf("a função Div retornou um valor inesperado para a entrada 1,1")
		}
	}
}
