package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	// Lee el txt
	content, err := ioutil.ReadFile(`ejercicio1.txt`)
	if err != nil {
		log.Fatal(err)
	}

	// Resive el contenido como una cadena
	text := string(content)

	// Cuenta la cantidad de caracteres
	charCount := len(text)

	// Cuenta la cantidad de palabras
	words := strings.Fields(text)
	wordCount := len(words)

	// Cuenta la cantidad de líneas
	lines := strings.Split(text, "\n")
	lineCount := len(lines)

	// Resta 2 al contador de caracteres por cada salto de línea
	charCount -= 2*lineCount - 2
	//Esto debido a que al contar los caracteres el "enter" se cuenta como
	//2 caracteres final de linal e inicio de la siguiente asi que al
	//contador de char se le restan 2 por cada linea, o sea por cada guion que encuentre
	//ademas recordar que el espacio, " ", es contado como un caracter

	// Imprime los resultados
	fmt.Printf("Cantidad de caracteres: %d\n", charCount)
	fmt.Printf("Cantidad de palabras: %d\n", wordCount)
	fmt.Printf("Cantidad de líneas: %d\n", lineCount)
}
