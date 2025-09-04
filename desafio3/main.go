package main

/*
=========== Desafio3 ===========
Escrever um código baseado nas aulas de concorrência, utilizando canais goroutines para que o programa exiba, em alternância, as palavras ping pong
*/
import (
	"fmt"
	"time"
)

// Função principal
func main() {
	// Criação dos canais c1 e c2, ambos do tipo string, para comunicação entre goroutines
	c1 := make(chan string)
	c2 := make(chan string)

	// Goroutine que envia "Pong" para o canal c1 após 1 segundo
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "Pong"
	}()

	// Goroutine que envia "Pong" para o canal c2 após 2 segundos
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "Pong"
	}()

	// Loop que aguarda mensagens dos canais e exibe "Ping Pong" conforme chegam
	for i := 0; i < 2; i++ {
		select { // Espera por uma mensagem de c1 ou c2 e executa o caso correspondente
		case msg1 := <-c1:
			fmt.Println("Ping", msg1)
		case msg2 := <-c2:
			fmt.Println("Ping", msg2)
		}
	}
}
