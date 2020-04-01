package main

import (
	"encoding/json"
	"net/http"
)

//Curso type
type Curso struct {
	//Para mostrar un display name en la respuesta del json
	Titulo       string
	NumeroVideos int
}

//Cursos array de cursos
type Cursos []Curso

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cursos := Cursos{
			Curso{"Curso GO", 30},
			Curso{"Curso C#", 30},
			Curso{"Curso Node JS", 30},
		}
		json.NewEncoder(w).Encode(cursos)
	})

	http.ListenAndServe(":8080", nil)
}
