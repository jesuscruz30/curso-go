// MANEJO DE ARCHIVOS

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Printf("D")
	leoArchivo()
	leoArchivo2()
	graboArchivo()
	graboArchivo2()
}

func leoArchivo() {
	archivo, err := ioutil.ReadFile("./archivo.txt")
	if err != nil {
		fmt.Println("Hubo un error")
	} else {
		fmt.Println(string(archivo))
	}
}

func leoArchivo2() {
	archivo, err := os.Open("./archivo.txt")
	if err != nil {
		fmt.Println("Hubo un error")
	} else {
		scanner := bufio.NewScanner(archivo)
		for scanner.Scan() {
			registro := scanner.Text()
			fmt.Printf("Linea >" + registro + "\n")
		}
	}
	archivo.Close()
}

func graboArchivo() {
	//crea el archivo o lo remplaza
	archivo, err := os.Create("./nuevoArchivo.txt")
	if err != nil {
		fmt.Println("Hubo un error")
		return
	}
	//Se graba en un archivo Fprintln
	fmt.Fprintln(archivo, "Esta es una linea nueva")
	archivo.Close()
}

func graboArchivo2() {
	//agrega lineas al contenido existente
	filename := "./nuevoArchivo.txt"
	if Append(filename, "\nEsta es una segunda linea") == false {
		fmt.Println("Error en la segunda linea")
	}
}

func Append(archivo string, texto string) bool {
	//el modo append no limpia el archivo
	arch, err := os.OpenFile(archivo, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Hubo un error")
		return false
	}
	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Hubo un error")
		return false
	}
	return true
}
