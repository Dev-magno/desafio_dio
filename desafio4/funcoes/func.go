package funcoes

// FUNÇÕES COM OPERAÇÕES ARITMÉTICAS

// Adição
func Adicao(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}

// Subtração
func Subtracao(i ...int) int {
	if len(i) == 0 {
		panic("Erro: Nenhum número fornecido")
	}
	total := i[0]
	for _, v := range i[1:] {
		total -= v
	}
	return total
}

// Divisão
func Divisao(i ...float64) float64 {
	if len(i) == 0 {
		panic("Erro: nenhum número fornecido")
	}
	total := i[0]
	for _, v := range i[1:] {
		if v == 0 {
			panic("Erro: Divisão por zero")
		}
		total /= v
	}
	return total
}

// Multiplicação
func Multiplicao(i ...int) int {
	total := 1
	for _, v := range i {
		total *= v
	}
	return total
}
