package main

import "fmt"

//[] se convierte automaticamente en un slice

func main() {
	slice := []int{1, 2, 3, 4}
	fmt.Println(slice)
	fmt.Println(len(slice))

	arreglo := [3]int{1, 2, 3} //Arreglo por que tiene tamañoo definico
	toSlice := arreglo[:2]     //Slicing
	fmt.Println(toSlice)
	toSlice = arreglo[1:2] //Slicing
	fmt.Println(toSlice)
	toSlice = arreglo[1:3] //Slicing
	fmt.Println(toSlice)
}

//Puntero al arreglo
//Longitud del arreglo al que apunta
//Capacidad
