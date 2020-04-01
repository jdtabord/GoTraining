package main

import (
	"fmt"

	"./dummy"
)

//Funciones que empiezan con minuscula son privadas
//Funciones que empiezan con mayuscula son publicas
func main() {
	fmt.Println(dummy.HolaMundo())
}
