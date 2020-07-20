//RECURSION
package main

import "fmt"

func main() {
	fmt.Println("")
	exponencia(2)
}

func exponencia(numero int) {
	if numero > 10000000000000000 {
		return
	}
	fmt.Println(numero)
	exponencia(numero * 2)
}
