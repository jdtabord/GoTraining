package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	executeReadFile()
	fmt.Println("Nunca me voy a imprimir")
}

func executeReadFile() {
	ejecucion := readFile()
	fmt.Println(ejecucion)
}

func readFile() bool {
	//Linea por linea
	archivo, error := os.Open("./hola.txt")

	//Fuciones que se tienen que ejecutar cuando la funcion actual retorna
	defer func() {
		archivo.Close()
		fmt.Println("Defer")
		r := recover()
		fmt.Println(r)
	}()

	if error != nil {
		panic(error)
	}

	scannner := bufio.NewScanner(archivo)
	var i int
	for scannner.Scan() {
		i++
		linea := scannner.Text()
		fmt.Println(i)
		fmt.Println(linea)
	}

	return true
}
