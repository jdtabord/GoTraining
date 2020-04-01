package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//Lee todo el archivo
	fileData, error := ioutil.ReadFile("./hola.txt")

	if error != nil {
		fmt.Println("Hubo un error")
	}

	fmt.Println(string(fileData))

	//Linea por linea
	archivo, error := os.Open("./hola.txt")

	//Fuciones que se tienen que ejecutar cuando la funcion actual retorna
	defer func() {
		archivo.Close()
		fmt.Println("Defer")
	}()

	if error != nil {
		fmt.Println("Hubo un error")
	}

	scannner := bufio.NewScanner(archivo)
	var i int
	for scannner.Scan() {
		i++
		linea := scannner.Text()
		fmt.Println(i)
		fmt.Println(linea)
	}

	archivo.Close()
}
