package main

import (
	"desafio4/funcoes"
	"fmt"
)

/*
=========== Desafio4 ===========
Criar uma calculadora com Go e executar testes unitários no código
*/
func main() {
	a := funcoes.Adicao(3, 6, 2)
	s := funcoes.Subtracao(8, 2)
	d := funcoes.Divisao(20, 5)
	m := funcoes.Multiplicao(10, 8)

	fmt.Println("Adição:", a)
	fmt.Println("Subtração: ", s)
	fmt.Println("Divisão: ", d)
	fmt.Println("Multiplicação:", m)
}
