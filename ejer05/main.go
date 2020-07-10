// CICLOS EN GO

package main

import "fmt"

func main() {
	/*for normal
	i := 1
	for i < 19 {
		fmt.Println(i)
		i++
	}
	*/
	//for con iteraciones definidas
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	//For infinito con condicion para break
	numero := 0
	for {
		fmt.Println("continuo")
		fmt.Println("ingresa el numero secreto")
		fmt.Scanln(&numero)
		if numero == 100 {
			break
		}
	}

	// for con condicion y continue
	//continue va al inicio de for
	var i = 0
	for i < 19 {
		fmt.Printf("\n Valor de i: %d", i)
		if i == 5 {
			fmt.Printf("multiplicamos por 2 \n")
			i = i * 2
			continue
		}
		fmt.Printf("Pasa por aqui \n")
		i++
	}
	// for con goto
	//goto va a una ubicacion especifica del codigo
	var x = 0
RUTINA:
	for x < 10 {
		if x == 4 {
			x = x + 2
			fmt.Println("voy a rutina sumando 2 a x")
			goto RUTINA
		}
		fmt.Printf("el valor de x: %d\n", x)
		x++
	}

}
