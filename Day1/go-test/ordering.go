package gotest

// função que ordena numeros inteiros de um slice e retorna o slice ordenado
func Order(numbers []int) []int {
	for i := 0; i < len(numbers); i++ {
		for j := i; j < len(numbers); j++ {
			if numbers[j] < numbers[i] {
				numbers[i], numbers[j] = numbers[j], numbers[i]
			}
		}
	}
	return numbers
}
