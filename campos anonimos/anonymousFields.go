package main

import "fmt"

//Human Type
type Human struct {
	nombre string
}

func (human Human) hablar() string {
	return "Bla bla bla"
}

//Dummy Type
type Dummy struct {
	nombre string
}

//Tutor Type
type Tutor struct {
	Human
	Dummy
}

func (tutor Tutor) hablar() string {
	return tutor.Human.hablar() + " Bienvenido Tutor"
}

func main() {
	tutor := Tutor{Human{"Juan"}, Dummy{"David"}}
	fmt.Println(tutor.Human.nombre)
	fmt.Println(tutor.Dummy.nombre)
	fmt.Println(tutor.hablar())
}
