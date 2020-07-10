//IF Y SWITCH

package main

import "fmt"

var estado bool

func main() {
	fmt.Println("=================IF==================")
	//se puede asignar el valor de la variable
	//en la misma comparacion del if
	estado = true
	if estado = false; estado == true {
		fmt.Println(estado)
	} else {
		fmt.Println("Es falso")
	}
	fmt.Println("=================SWITCH==================")
	//switch
	switch numero := 5; numero {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)
	case 5:
		fmt.Println(5)
	default:
		fmt.Println("Mayor a 5")
	}

}
