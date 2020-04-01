package main

import "fmt"

//Crear un slice
//Make(tipo, longitud, capacidad)
//len tamaño del arreglo
//cap cuantos elementos caben en un slice
//Puntero direccion en memoria

func main() {
	slice := make([]int, 3, 5)
	slice = append(slice, 2)
	fmt.Println(slice)
}
