//Channels en GO (Dialogo entre GORoutines)
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("")
	canal1 := make(chan time.Duration)
	go bucle(canal1)
	fmt.Println("Llegue hasta aqui")
	//await nodejs, promesa
	//espera hasta que msg tenga un valor
	msg := <-canal1
	fmt.Println(msg)
}

func bucle(canal1 chan time.Duration) {
	inicio := time.Now()
	for i := 0; i < 100000000000; i++ {

	}
	final := time.Now()
	//sub es la funcion que retorna la duración
	canal1 <- final.Sub(inicio)
}
