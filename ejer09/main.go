//MAPAS EN GO
/* una estructura que se puede serializar
en un mapa el indice es clave y valor asociado a la clave
*/
package main

import "fmt"

func main() {
	// [string] es la clave
	// string es el valor
	paises := make(map[string]string)
	//paises2 reserva espacio para 5 elementos
	//paises2 := make(map[string]string, 5)
	fmt.Println(paises)
	paises["Mexico"] = "D.F"
	paises["Argentina"] = "Buenos Aires"
	fmt.Println(paises["Mexico"])

	fmt.Println("Imperimir campeonato :")
	campeonato := map[string]int{
		"Barcelona":   39,
		"Real Madrid": 38,
		"Chivas":      37,
		"Boca junior": 30}
	fmt.Println(campeonato)
	//agregar un nuevo dato
	campeonato["River Plate"] = 25
	//cambiar un valor
	campeonato["Chivas"] = 55
	fmt.Println(campeonato)
	//borrar un valor
	delete(campeonato, "Real Madrid")
	fmt.Println(campeonato)

	for equipo, puntaje := range campeonato {
		fmt.Printf("El equipo %s,tiene un puntaje de : %d \n", equipo, puntaje)
	}

	//verificar si un elemento existe en un mapa
	//nota %t es para booleano
	puntaje, existe := campeonato["Chivas"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe %t \n", puntaje, existe)
}
