package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4}
	copia := make([]int, len(slice))

	//Copy copia el minimo de elementos con base en la longitud del destino
	//copy(destino, fuente)
	copy(copia, slice)

	fmt.Println(slice)
	fmt.Println(copia)

}
