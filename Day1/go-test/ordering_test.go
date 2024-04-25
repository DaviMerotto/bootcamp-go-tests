package gotest

import "testing"

// função de Test da função Order
func TestOrder(t *testing.T) {
	// cria um slice de inteiros
	numbers := []int{6, 4, 45}
	// chama a função Order
	ordered := Order(numbers)
	// verifica se o slice retornado está ordenado
	if ordered[0] != 4 || ordered[1] != 6 || ordered[2] != 45 {
		t.Errorf("a função Order não ordenou o slice corretamente")
	}
}
