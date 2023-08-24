package main

import (
	"fmt"
)

func rotateSequence(sequence []string, rotations int, direction int) []string {
	length := len(sequence)
	result := make([]string, length)

	for i := 0; i < length; i++ {
		var newIndex int
		if direction == 0 {
			newIndex = (i + rotations) % length
		} else if direction == 1 {
			newIndex = (i - rotations + length) % length
		}
		result[newIndex] = sequence[i]
	}

	return result
}

func main() {
	originalSequence := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	rotations := 3
	direction := 1

	fmt.Println("Secuencia Original =", originalSequence)
	fmt.Println("Cantidad de rotaciones =", rotations)
	if direction == 0 {
		fmt.Println("Dirección = izq")
	} else if direction == 1 {
		fmt.Println("Dirección = der")
	}

	rotatedSequence := rotateSequence(originalSequence, rotations, direction)
	fmt.Println("Secuencia final rotada =", rotatedSequence)
}
