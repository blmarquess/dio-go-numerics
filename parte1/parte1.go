package parte1

import (
	"fmt"
)

// Desafio, criar um programa que imprima os números de 1 a 1000 utilizando a estrutura de repetição for.
// Mais podemos imprimir apenas os números divisíveis por 3.

func PartOne() {
	for i := 1; i <= 1000; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}
