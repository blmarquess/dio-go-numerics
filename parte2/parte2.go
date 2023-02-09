package parte2

import (
	"fmt"
)

// Desafio escrever um programa que ao encontrar um numero múltiplo de 3 escreve no terminal pin e se for múltiplo de 5 escreve pan e se for múltiplo de 3 e 5 escreve pinpan

func PinPan() {
	for i := 0; i < 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("PinPan")
			continue
		}
		if i%3 == 0 {
			fmt.Println("Pin")
			continue
		}
		if i%5 == 0 {
			fmt.Println("Pan")
			continue
		}
	}
}
