package main

import (
	"fmt"
)

func main() {
	channel := make(chan string)

	go func(channel chan string) {
		for {
			var name string
			fmt.Scanln(&name)
			channel <- name
		}
	}(channel)

	//Espera a que el canal reciba informacion
	msg := <-channel

	fmt.Println("Estoy imprimiendo lo que recibi del canal: " + msg)

	msg = <-channel

	fmt.Println("Estoy imprimiendo lo que recibi del canal: " + msg)

}
