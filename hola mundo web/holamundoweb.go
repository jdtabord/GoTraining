package main

import (
	"net/http"
)

func main() {
	//cuando ingrese al 8080 con barra se va a la funcion handler
	fileServer := http.FileServer(http.Dir("publico"))
	//Para proteger los demas archivos
	http.Handle("/public/", http.StripPrefix("/public", fileServer))
	http.ListenAndServe(":8080", nil)
}

//ResponseWriter para responder la peticion
//Request puntero a la informacion de la peticio(que nos envia el navegador)
// func handler(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w, r, r.URL.Path[1:])
// }
