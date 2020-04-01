package main

import "fmt"

func main() {
	/*
		1. Es una direccion de memoria
		2. Todas las variables estan almacenadas en en una direccion
		3. Con un puntero conocemos la direccion y no el valor
		4. X,Y ===> as123d => 5
		5. x ===> as123d => 6
		6. Y => 6
		7. Esto por que ambas variables estan en la misma posicion de memoria
		8. El valor 0 en puntero es nil(int o, bool false, string vacio)
	*/

	var X, Y *int
	entero := 5

	//Con &variable obtenemos la direccion en memoria
	X = &entero
	Y = &entero

	fmt.Println(X)
	fmt.Println(Y)

	*X = 6

	//Con *variable obtenemos el valor
	fmt.Println(*X)
	fmt.Println(*Y)

}
