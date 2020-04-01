package main

import "fmt"

//User Interface
type User interface {
	Permisos() int
	Nombre() string
}

//Admin type
type Admin struct {
	nombre string
}

//Permisos metodo para retornar los permisos
func (admin Admin) Permisos() int {
	return 5
}

//Nombre metodo para devolver el nombre
func (admin Admin) Nombre() string {
	return admin.nombre
}

//Auth metodo para autenticar
func Auth(user User) string {
	permisos := user.Permisos()
	if permisos > 4 {
		return user.Nombre() + " tiene permisos de administrador"
	}
	return "no tiene permisos de administrador"
}

func main() {
	admin := Admin{nombre: "Juan"}
	fmt.Println(Auth(admin))
}
