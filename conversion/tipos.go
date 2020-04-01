package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Itoa = int a string
	edad := 22
	edadStr := strconv.Itoa(edad)
	fmt.Println("Mi edad es " + edadStr)

	//Atoi = string to int
	estrato := "3"
	estratoInt, _ := strconv.Atoi(estrato)
	fmt.Println(estratoInt)
}
