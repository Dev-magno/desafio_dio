package main

import "fmt"

/*      =============== DESAFIO ===============
Criar código que exiba todos os números divisíveis por 3 de 1 a 100
*/

func DivisivelPorTres() {
	for d := 1; d <= 100; d++ {
		if d%3 == 0 {
			fmt.Println("Divisíveis por 3: ", d)
		}
	}
}

/*      =============== DESAFIO - parte 2 ===============
Criar um programa que exiba a palavra "Pin" quando for múltiplo de 3 e "Pan" quando for múltiplo de 5. Se for multiplo de 3 e 5 ao mesmo tempo, escrever "PinPan".
Obs: Os números não devem aparecer
*/

func FalarPinPan() {
	for p := 1; p <= 100; p++ {
		if p%3 == 0 && p%5 == 0 {
			fmt.Println("PinPan")
		} else if p%3 == 0 {
			fmt.Println("Pin")
		} else if p%5 == 0 {
			fmt.Println("Pan")
		}
	}
}
func main() {
	DivisivelPorTres()
	fmt.Println("====== Retorna Pin Pan ======")
	FalarPinPan()
}
