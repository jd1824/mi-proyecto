package main

import "fmt"

func invertir(cadena string) string {
	runas := []rune(cadena)
	longitud := len(runas)
	salida := make([]rune, longitud)

	for i, r := range runas {
		salida[longitud-1-i] = r
	}

	return string(salida)
}

func main() {
	fmt.Println(invertir("aibofobia "))
}
