//FUNCIONES
package main

import "fmt"

func main() {
	fmt.Println("======Ejecucion funcion 1=====")
	fmt.Println(uno(5))
	//crea dos variables del resultado de los
	//dos valores que devuelve dos
	numero, estado := dos(1)
	fmt.Println("======Ejecucion funcion dos====")
	fmt.Println(numero)
	fmt.Println(estado)
	fmt.Println("======Ejecucion funcion calculo====")
	fmt.Println(calculo(5, 46))
	fmt.Println(calculo(2, 23, 45, 68))
	fmt.Println(calculo(5))
	fmt.Println(calculo(5, 46, 17, 25, 26, 98, 17))

}

// RETORNO DE UN PARAMETRO
func uno(numero int) (z int) {
	z = numero * 2
	return
}

//RETORNO DE DOS PARAMETROS RESPUESTA VARIABLE
func dos(numero int) (int, bool) {
	if numero == 1 {
		return 5, false
	} else {
		return 10, true
	}
}

// nota: en Go no existe sobrecarga de metodos

// PARAMETROS DE ENTRADA VARIABLES
func calculo(numero ...int) int {
	total := 0
	//convierte los parametros recibidis
	//en una lista
	//range siempre devuelve dos valores
	//el primero es el numero de elementos
	//por eso se usa guion bajo _ , en los casos donde se aloja
	//la variable que no se va a usar

	//for _, num := range numero {
	for item, num := range numero {
		total = total + num
		fmt.Printf("\n Item : %d \n", item)
	}
	return total
}
