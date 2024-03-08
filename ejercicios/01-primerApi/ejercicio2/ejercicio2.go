package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Persona struct {
	Nombre   string `json:"firstName"`
	Apellido string `json:"lastName"`
}

func main() {
	router := chi.NewRouter()

	router.Post("/greetings", func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var persona Persona

		err := decoder.Decode(&persona)
		if err != nil {
			w.Write([]byte("Error: " + err.Error()))
			return
		}

		w.Write([]byte("Hello " + persona.Nombre + " " + persona.Apellido))
	})

	http.ListenAndServe(":8080", router)
}
