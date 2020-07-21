//Middlewares en GO
package main

import "fmt"

var result int

func main() {
	fmt.Println("Inicio")
	//middleware operacionesMidd
	result = operacionesMidd(sumar)(2, 3)
	fmt.Println(result)

	result = operacionesMidd(restar)(2, 10)
	fmt.Println(result)

	result = operacionesMidd(multiplicar)(2, 4)
	fmt.Println(result)
}

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

//cuando un parametro es un afuncion y esa funcion
//recibe valores enteros no se pone el nombre solo el tipo
//lo que recibe el middleware y lo que devuelve tiene que ser igual
func operacionesMidd(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Inicio de operacion")
		return f(a, b)
	}

}
