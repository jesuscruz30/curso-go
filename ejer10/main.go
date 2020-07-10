// ESTRUCTURAS POO EN GO
package main

//us hace referencia a user
import (
	"fmt"
	"time"

	us "./user"
)

type pepe struct {
	us.Usuario
}

func main() {
	u := new(pepe)
	u.AltaUsuario(1, "Pablo Tilotta", time.Now(), true)
	fmt.Println(u.Usuario)
}

//se remplaza esto con user.go
/*type usuario struct {
	Id        int
	Nombre    string
	FechaAlta time.Time
	Status    bool
}


func main() {
	User := new(usuario)
	User.Id = 10
	User.Nombre = "Pablo"
	fmt.Println(User)
}
*/
