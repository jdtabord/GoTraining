package main

import "fmt"

//[10] se establece el tamaño y no se puede modificar
//[3] int {1,2,3} se declara con valores

func main() {
	var arregloNumeros [10]int
	fmt.Println(arregloNumeros)

	arregloLetras := [3]string{"a", "b", "c"}
	fmt.Println(arregloLetras)

	for i := 0; i < len(arregloNumeros); i++ {
		arregloNumeros[i] = i
	}

	fmt.Println(arregloNumeros)

	var matriz [2][2]int
	fmt.Println(matriz)

}
