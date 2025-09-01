package main

import "fmt"

// Ponto de ebulição
const ebulicaoK = 373.15

func main() {

	tempC := ebulicaoK - 273.15
	fmt.Printf("A temperatura de ebulição da água em Celsius é: %g°C\n", tempC)
}
