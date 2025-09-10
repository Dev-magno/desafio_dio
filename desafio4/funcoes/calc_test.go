package funcoes_test

import (
	"desafio4/funcoes"
	"testing"
)

func TestAdd(t *testing.T) {
	teste := funcoes.Adicao(2, 3, 6)
	resultado := 11

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornado: ", teste)
	}
}

func TestSubtract(t *testing.T) {
	teste := funcoes.Subtracao(8, 2)
	resultado := 6

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornado: ", teste)
	}
}

func TestSubtracaoEmptyInput(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for empty input, but none occurred")
		}
	}()
	funcoes.Subtracao() // Deve causar panic
}

func TestDivide(t *testing.T) {
	teste := funcoes.Divisao(20, 5)
	resultado := 4

	if teste != float64(resultado) {
		t.Error("Valor esperado: ", resultado, "Valor retornado: ", teste)
	}
}

func TestDivideEmptyInput(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Era esperado um pânico para entrada vazia, mas nenhum ocorreu.")
		}
	}()
	funcoes.Divisao()
}

func TestDivideByZero(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Esperava-se um panic por divisão por zero, mas isso não aconteceu.")
		}
	}()
	funcoes.Divisao(10.0, 0.0)
}

func TestMultiply(t *testing.T) {
	teste := funcoes.Multiplicao(10, 8)
	resultado := 80

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornado: ", teste)
	}
}
