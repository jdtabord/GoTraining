package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	go miNombreLento("JuanDavid")
	fmt.Println("Puta Cuarentena")
	var wait string
	fmt.Scanln(&wait)
}

func miNombreLento(name string) {
	//"" en split separa letra a letra
	letras := strings.Split(name, "")
	//_ para ignorar el index
	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
}
