package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//Todo en la misma linea
	fmt.Print("Hola Mundo")

	//Hace el salto de linea
	fmt.Println("Hola Mundo con salto de linea")

	//Para concatenar datos
	//%v para cualquier tipo de dato
	edad := 24
	bandera := true
	fmt.Printf("Mi edad es %v\n", edad)
	fmt.Printf("Veradero es %v\n", bandera)

	var edad2 int
	fmt.Println("Ingresa tu edad: ")
	//Scanf para leer datos
	//%d para leer el dato entero
	fmt.Scanf("%d\n", &edad2)
	fmt.Printf("Tu edad es %d\n", edad2)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ingresa tu nombre: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Hola " + text)
	}

}
