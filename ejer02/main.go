//VARIABLES

package main

import (
	"fmt"
	"strconv"
)

var numero int
var texto string
var status bool = true

func main() {
	numero2, numero3, numero4, texto, status := 2, 5, 6, "Este s mi texto", false
	var moneda int64 = 1

	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println(numero4)
	fmt.Println(texto)
	fmt.Println(status)
	numero2 = numero2 + int(moneda)
	fmt.Println("nuevo valor de numero2")
	fmt.Println(numero2)

	texto = strconv.Itoa(int(moneda))
	fmt.Println(texto)
	mostrarStatus()
}

func mostrarStatus() {
	fmt.Println(status)
}
