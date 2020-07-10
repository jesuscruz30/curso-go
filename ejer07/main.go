// FUNCIONES ANONIMAS Y CLOSURES
package main

import "fmt"

//FUNCION ANONIMA calculo
var calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Printf("Sumo 5 + 7 = %d \n", calculo(5, 7))

	calculo = func(num1 int, num2 int) int {
		return num1 - num2
	}
	fmt.Printf("Resta 6 - 4 = %d \n", calculo(6, 4))

	calculo = func(num1 int, num2 int) int {
		return num1 / num2
	}
	fmt.Printf("Division 12 / 3 = %d \n", calculo(12, 3))

	operaciones()
	/*Closures*/
	fmt.Println("Tabla del 2 :")
	tablDel := 2
	miTabla := tabla(tablDel)
	for i := 1; i < 11; i++ {
		fmt.Println(miTabla())
	}

}

func operaciones() {
	resultado := func() int {
		var a int = 23
		var b int = 13
		return a + b
	}
	fmt.Println(resultado())
}

/* CLOSURES isolacion de codigo
 para proteger variables creadas en la funcion original

la funcion tabla devuelve otra funcion de tipo int
*/

func tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia++
		return numero * secuencia
	}
}
