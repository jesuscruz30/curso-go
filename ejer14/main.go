//DEFER PANIC & RECOVER
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//ejemploDefer()
	ejemploPanic()
}

func ejemploDefer() {
	//va a arrojar un error
	archivo := "prueba.txt"
	f, err := os.Open(archivo)

	//el defer ejecuta la instruccion antes de salir de la funcion

	defer f.Close()

	if err != nil {
		fmt.Println("Error abriendo el archivo")
		os.Exit(1)
	}

}

func ejemploPanic() {
	//funcion anonima
	defer func() {
		reco := recover()
		if reco != nil {
			//fatalF graba en log y es similar a printF
			log.Fatalf("Ocurrio un error que generó Panic \n %v", reco)
		}
	}()
	//se ponen parentesis al final de la funcion anonima porque
	//no regresara nada
	a := 1
	if a == 1 {
		panic("se encontro el valor de 1")
	}
}
