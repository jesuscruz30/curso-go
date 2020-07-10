/*MATRICES, ARRAYS Y SLICES */
package main

import "fmt"

var tabla [10]int
var matriz [5][7]int

//tabla2 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
func main() {
	tabla[0] = 1
	tabla[5] = 15
	fmt.Println(tabla)

	//vector completo
	tabla2 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("\ntabla2 :", tabla2)

	//recorrer vector con for
	fmt.Println("\nTabla recorrida con for")
	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}

	//matriz completa
	fmt.Println("\n matriz ")
	matriz[3][5] = 1
	fmt.Println(matriz)

	//slices
	fmt.Println("\n matriz slice")
	matrizSlice := []int{2, 5, 4}
	fmt.Println(matrizSlice)
	variante2()
	variante3()
	variante4()
}

func variante2() {
	fmt.Println("variante2 de slice")
	elementos := [5]int{1, 2, 3, 4, 5}
	//crear un slice en base a un vextor
	porcion := elementos[3:]
	fmt.Println(porcion)
}

func variante3() {
	fmt.Println("\n variante3 de slice ")
	//crea en memoria un espacio con 20 elementos
	//pero inicialmente crea un vector de 5 elementos
	elementos := make([]int, 5, 20)
	//len longitud
	//cap capacidad
	//nota : Printf permite mezclar texto y datos
	fmt.Printf("Largo %d, capacidad %d \n", len(elementos), cap(elementos))

}

func variante4() {
	fmt.Println("\n variante4 de slice ")
	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	//da capacidad de 128 , segun el siguiente
	// a los numeros binarios 2,4,8,16,64,128
	//como en este caso es 100 el siguiente mas cercano
	//es 100 -> 128
	fmt.Printf("\n Largo %d, Capacidad %d", len(nums), cap(nums))
}
