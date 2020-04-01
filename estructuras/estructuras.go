package main

import "fmt"

//User Type indica que se crea un nuevo tipo que es una estructura
type User struct {
	edad             int
	nombre, apellido string
}

func main() {
	user := new(User)
	user.edad = 24
	user.nombre = "Juan"
	user.apellido = "Taborda"
	fmt.Println(user.nombre)
}
