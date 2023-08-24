package main

import (
	"fmt"
)

func main() {
	ancho := 5

	if ancho%2 == 0 {
		fmt.Println("El ancho debe ser un número impar.")
		return
	}

	// Imprimir la mitad superior del rombo
	for i := 1; i <= ancho; i += 2 {
		for j := 0; j < (ancho-i)/2; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// Imprimir la mitad inferior del rombo
	for i := ancho - 2; i >= 1; i -= 2 {
		for j := 0; j < (ancho-i)/2; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
