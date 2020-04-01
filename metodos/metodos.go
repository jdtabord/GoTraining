package main

import "fmt"

//User type for users
type User struct {
	nombre, apellido string
}

func (user User) nombreCompleto() string {
	return user.nombre + " " + user.apellido
}

//Se pasa la referencia como puntero para poder hacer una copia y modificar su valor
func (user *User) setName(n string) {
	user.nombre = n
}

func main() {
	usuario := new(User)
	usuario.nombre = "Juan"
	usuario.apellido = "Taborda"

	fmt.Println(usuario.nombreCompleto())

	usuario.setName("David")
	fmt.Println(usuario.nombre)

}
